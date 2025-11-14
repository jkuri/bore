package server

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"go.uber.org/zap"
)

type TunnelMetrics struct {
	ID                 string    `json:"id"`
	Domain             string    `json:"domain"`
	Port               uint32    `json:"port"`
	Addr               string    `json:"addr"`
	BytesIn            uint64    `json:"bytesIn"`
	BytesOut           uint64    `json:"bytesOut"`
	CumulativeBytesIn  uint64    `json:"cumulativeBytesIn"`
	CumulativeBytesOut uint64    `json:"cumulativeBytesOut"`
	ThroughputIn       float64   `json:"throughputIn"`
	ThroughputOut      float64   `json:"throughputOut"`
	ConnectedAt        time.Time `json:"connectedAt"`
	LastActivity       time.Time `json:"lastActivity"`
	ActiveConnections  int       `json:"activeConnections"`
}

type ServerStats struct {
	TotalBytesIn       uint64  `json:"totalBytesIn"`
	TotalBytesOut      uint64  `json:"totalBytesOut"`
	CumulativeBytesIn  uint64  `json:"cumulativeBytesIn"`
	CumulativeBytesOut uint64  `json:"cumulativeBytesOut"`
	ThroughputIn       float64 `json:"throughputIn"`
	ThroughputOut      float64 `json:"throughputOut"`
}

type DashboardMessage struct {
	Type        string          `json:"type"`
	Tunnels     []TunnelMetrics `json:"tunnels"`
	ServerStats ServerStats     `json:"serverStats"`
}

type MetricsHub struct {
	mu              sync.RWMutex
	clients         map[*websocket.Conn]bool
	tunnelMetrics   map[string]*TunnelMetrics
	sshServer       *SSHServer
	logger          *zap.SugaredLogger
	broadcastTicker *time.Ticker
	serverStats     ServerStats
	lastStatsUpdate time.Time
}

func NewMetricsHub(sshServer *SSHServer, logger *zap.SugaredLogger) *MetricsHub {
	hub := &MetricsHub{
		clients:         make(map[*websocket.Conn]bool),
		tunnelMetrics:   make(map[string]*TunnelMetrics),
		sshServer:       sshServer,
		logger:          logger,
		lastStatsUpdate: time.Now(),
	}

	hub.broadcastTicker = time.NewTicker(200 * time.Millisecond)
	go hub.broadcastLoop()

	return hub
}

func (h *MetricsHub) broadcastLoop() {
	for range h.broadcastTicker.C {
		h.broadcastMetrics()
	}
}

func (h *MetricsHub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		h.logger.Errorf("websocket accept error: %v", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()

	h.logger.Infof("new websocket client connected")

	h.sendInitialMetrics(conn)

	ctx := r.Context()
	for {
		_, _, err := conn.Read(ctx)
		if err != nil {
			break
		}
	}

	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()

	h.logger.Infof("websocket client disconnected")
}

func (h *MetricsHub) sendInitialMetrics(conn *websocket.Conn) {
	metrics := h.collectMetrics()

	h.mu.RLock()
	serverStats := h.serverStats
	h.mu.RUnlock()

	msg := DashboardMessage{
		Type:        "update",
		Tunnels:     metrics,
		ServerStats: serverStats,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		h.logger.Errorf("failed to marshal metrics: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
		h.logger.Errorf("failed to send initial metrics: %v", err)
	}
}

func (h *MetricsHub) broadcastMetrics() {
	metrics := h.collectMetrics()
	h.updateServerStats()

	h.mu.RLock()
	serverStats := h.serverStats
	clients := make([]*websocket.Conn, 0, len(h.clients))
	for client := range h.clients {
		clients = append(clients, client)
	}
	h.mu.RUnlock()

	msg := DashboardMessage{
		Type:        "update",
		Tunnels:     metrics,
		ServerStats: serverStats,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		h.logger.Errorf("failed to marshal metrics: %v", err)
		return
	}

	for _, client := range clients {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := client.Write(ctx, websocket.MessageText, data); err != nil {
			h.logger.Errorf("failed to broadcast to client: %v", err)
			h.mu.Lock()
			delete(h.clients, client)
			h.mu.Unlock()
		}
		cancel()
	}
}

func (h *MetricsHub) collectMetrics() []TunnelMetrics {
	h.sshServer.mu.Lock()
	defer h.sshServer.mu.Unlock()

	metrics := make([]TunnelMetrics, 0, len(h.sshServer.clients))
	now := time.Now()

	for id, client := range h.sshServer.clients {
		h.mu.Lock()
		metric, exists := h.tunnelMetrics[id]
		if !exists {
			metric = &TunnelMetrics{
				ID:           id,
				Domain:       h.sshServer.domain,
				Port:         client.port,
				Addr:         client.addr,
				ConnectedAt:  now,
				LastActivity: now,
			}
			h.tunnelMetrics[id] = metric
		}

		metric.Port = client.port
		metric.Addr = client.addr

		elapsed := now.Sub(metric.LastActivity).Seconds()
		if elapsed >= 1.0 {
			metric.ThroughputIn = float64(metric.BytesIn) / elapsed
			metric.ThroughputOut = float64(metric.BytesOut) / elapsed
			metric.BytesIn = 0
			metric.BytesOut = 0
			metric.LastActivity = now
		}

		client.mu.Lock()
		metric.ActiveConnections = len(client.channels)
		client.mu.Unlock()

		metrics = append(metrics, *metric)
		h.mu.Unlock()
	}

	h.mu.Lock()
	for id := range h.tunnelMetrics {
		if _, exists := h.sshServer.clients[id]; !exists {
			delete(h.tunnelMetrics, id)
		}
	}
	h.mu.Unlock()

	return metrics
}

func (h *MetricsHub) RecordTraffic(tunnelID string, bytesIn, bytesOut uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if metric, exists := h.tunnelMetrics[tunnelID]; exists {
		metric.BytesIn += bytesIn
		metric.BytesOut += bytesOut
		metric.CumulativeBytesIn += bytesIn
		metric.CumulativeBytesOut += bytesOut
		metric.LastActivity = time.Now()
	}

	h.serverStats.TotalBytesIn += bytesIn
	h.serverStats.TotalBytesOut += bytesOut
	h.serverStats.CumulativeBytesIn += bytesIn
	h.serverStats.CumulativeBytesOut += bytesOut
}

func (h *MetricsHub) updateServerStats() {
	h.mu.Lock()
	defer h.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(h.lastStatsUpdate).Seconds()

	if elapsed >= 1.0 {
		h.serverStats.ThroughputIn = float64(h.serverStats.TotalBytesIn) / elapsed
		h.serverStats.ThroughputOut = float64(h.serverStats.TotalBytesOut) / elapsed
		h.serverStats.TotalBytesIn = 0
		h.serverStats.TotalBytesOut = 0
		h.lastStatsUpdate = now
	}
}

func (h *MetricsHub) Close() {
	if h.broadcastTicker != nil {
		h.broadcastTicker.Stop()
	}
}
