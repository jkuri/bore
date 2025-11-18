import { Activity, ArrowDown, ArrowUp, Network } from "lucide-react";
import { useEffect, useState } from "react";
import { Badge } from "./ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";

interface TunnelMetrics {
  id: string;
  domain: string;
  port: number;
  addr: string;
  bytesIn: number;
  bytesOut: number;
  cumulativeBytesIn: number;
  cumulativeBytesOut: number;
  throughputIn: number;
  throughputOut: number;
  connectedAt: string;
  lastActivity: string;
  activeConnections: number;
}

interface ServerStats {
  totalBytesIn: number;
  totalBytesOut: number;
  cumulativeBytesIn: number;
  cumulativeBytesOut: number;
  throughputIn: number;
  throughputOut: number;
}

interface DashboardMessage {
  type: string;
  tunnels: TunnelMetrics[];
  serverStats: ServerStats;
}

function formatBytes(bytes: number): string {
  if (bytes === 0) return "0 B/s";
  const k = 1024;
  const sizes = ["B/s", "KB/s", "MB/s", "GB/s"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / k ** i).toFixed(2)} ${sizes[i]}`;
}

function formatDuration(dateStr: string): string {
  const date = new Date(dateStr);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const seconds = Math.floor(diff / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (days > 0) return `${days}d ${hours % 24}h`;
  if (hours > 0) return `${hours}h ${minutes % 60}m`;
  if (minutes > 0) return `${minutes}m ${seconds % 60}s`;
  return `${seconds}s`;
}

function formatTotalBytes(bytes: number): string {
  if (bytes === 0) return "0 B";
  const k = 1024;
  const sizes = ["B", "KB", "MB", "GB", "TB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${(bytes / k ** i).toFixed(2)} ${sizes[i]}`;
}

export default function Dashboard(): React.JSX.Element {
  const [tunnels, setTunnels] = useState<TunnelMetrics[]>([]);
  const [serverStats, setServerStats] = useState<ServerStats>({
    totalBytesIn: 0,
    totalBytesOut: 0,
    cumulativeBytesIn: 0,
    cumulativeBytesOut: 0,
    throughputIn: 0,
    throughputOut: 0,
  });
  const [connected, setConnected] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
    const wsUrl = `${protocol}//${window.location.host}/api/ws/dashboard`;

    let ws: WebSocket | null = null;
    let reconnectTimeout: ReturnType<typeof setTimeout>;

    const connect = () => {
      try {
        ws = new WebSocket(wsUrl);

        ws.onopen = () => {
          setConnected(true);
          setError(null);
        };

        ws.onmessage = (event) => {
          try {
            const data: DashboardMessage = JSON.parse(event.data);
            if (data.type === "update") {
              const tunnels = data.tunnels.sort((a, b) => {
                return a.id.localeCompare(b.id);
              });
              setTunnels(tunnels || []);
              setServerStats(
                data.serverStats || {
                  totalBytesIn: 0,
                  totalBytesOut: 0,
                  cumulativeBytesIn: 0,
                  cumulativeBytesOut: 0,
                  throughputIn: 0,
                  throughputOut: 0,
                }
              );
            }
          } catch (err) {
            console.error("Failed to parse message:", err);
          }
        };

        ws.onerror = () => {
          setError("WebSocket connection error");
          setConnected(false);
        };

        ws.onclose = () => {
          setConnected(false);
          reconnectTimeout = setTimeout(connect, 3000);
        };
      } catch (_err) {
        setError("Failed to connect to WebSocket");
        setConnected(false);
        reconnectTimeout = setTimeout(connect, 3000);
      }
    };

    connect();

    return () => {
      if (ws) {
        ws.close();
      }
      if (reconnectTimeout) {
        clearTimeout(reconnectTimeout);
      }
    };
  }, []);

  const totalConnections = tunnels.reduce(
    (sum, t) => sum + t.activeConnections,
    0
  );

  return (
    <div className="mx-auto max-w-7xl space-y-6 p-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="font-semibold text-3xl tracking-tight">Dashboard</h1>
          <p className="text-muted-foreground text-sm">
            Real-time tunnel monitoring
          </p>
        </div>
        {connected ? (
          <Badge variant="outline" className="gap-1.5">
            <span className="h-2 w-2 animate-pulse rounded-full bg-green-500 dark:bg-green-400" />
            Live
          </Badge>
        ) : (
          <Badge variant="outline" className="gap-1.5">
            <span className="h-2 w-2 rounded-full bg-muted-foreground/50" />
            Disconnected
          </Badge>
        )}
      </div>

      {error && (
        <div className="rounded-lg border bg-card p-3 text-card-foreground text-sm">
          {error}
        </div>
      )}

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">
              Active Tunnels
            </CardTitle>
            <Network className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">{tunnels.length}</div>
            <p className="text-muted-foreground text-xs">
              {totalConnections} active channel
              {totalConnections !== 1 ? "s" : ""}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">Throughput In</CardTitle>
            <ArrowDown className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">
              {formatBytes(serverStats.throughputIn)}
            </div>
            <p className="text-muted-foreground text-xs">
              {formatTotalBytes(serverStats.cumulativeBytesIn)} total
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">
              Throughput Out
            </CardTitle>
            <ArrowUp className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">
              {formatBytes(serverStats.throughputOut)}
            </div>
            <p className="text-muted-foreground text-xs">
              {formatTotalBytes(serverStats.cumulativeBytesOut)} total
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="font-medium text-sm">Total Traffic</CardTitle>
            <Activity className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="font-bold text-2xl">
              {formatBytes(
                serverStats.throughputIn + serverStats.throughputOut
              )}
            </div>
            <p className="text-muted-foreground text-xs">
              {formatTotalBytes(
                serverStats.cumulativeBytesIn + serverStats.cumulativeBytesOut
              )}{" "}
              total
            </p>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Tunnels</CardTitle>
          <CardDescription>
            Monitor all active tunnel connections and their traffic
          </CardDescription>
        </CardHeader>
        <CardContent>
          {tunnels.length === 0 ? (
            <div className="flex flex-col items-center justify-center py-10 text-center">
              <div className="mb-4 rounded-full bg-muted p-3">
                <Network className="h-6 w-6 text-muted-foreground" />
              </div>
              <h3 className="mb-1 font-semibold text-sm">No active tunnels</h3>
              <p className="text-muted-foreground text-sm">
                Tunnels will appear here when clients connect
              </p>
            </div>
          ) : (
            <div className="rounded-md border">
              <Table>
                <TableHeader>
                  <TableRow className="hover:bg-transparent">
                    <TableHead className="h-10 text-xs">ID</TableHead>
                    <TableHead className="h-10 text-xs">URL</TableHead>
                    <TableHead className="h-10 text-xs">Port</TableHead>
                    <TableHead className="h-10 text-xs">Uptime</TableHead>
                    <TableHead className="h-10 text-xs">Channels</TableHead>
                    <TableHead className="h-10 w-36 text-right text-xs">
                      In
                    </TableHead>
                    <TableHead className="h-10 w-36 text-right text-xs">
                      Out
                    </TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {tunnels.map((tunnel) => (
                    <TableRow key={tunnel.id} className="h-12">
                      <TableCell className="py-1 font-medium font-mono text-xs">
                        {tunnel.id}
                      </TableCell>
                      <TableCell className="py-1 text-xs">
                        <a
                          href={`http://${tunnel.id}.${tunnel.domain}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-primary hover:underline"
                        >
                          {tunnel.id}.{tunnel.domain}
                        </a>
                      </TableCell>
                      <TableCell className="py-1 font-mono text-muted-foreground text-xs">
                        {tunnel.port}
                      </TableCell>
                      <TableCell className="py-1 text-muted-foreground text-xs">
                        {formatDuration(tunnel.connectedAt)}
                      </TableCell>
                      <TableCell className="py-1">
                        <Badge
                          variant={
                            tunnel.activeConnections > 0
                              ? "default"
                              : "secondary"
                          }
                          className="h-5 px-1.5 text-xs"
                        >
                          {tunnel.activeConnections}
                        </Badge>
                      </TableCell>
                      <TableCell className="py-1 text-right">
                        <div className="font-mono text-xs tabular-nums">
                          {formatBytes(tunnel.throughputIn)}
                        </div>
                        <div className="text-[10px] text-muted-foreground">
                          {formatTotalBytes(tunnel.cumulativeBytesIn)} total
                        </div>
                      </TableCell>
                      <TableCell className="py-1 text-right">
                        <div className="font-mono text-xs tabular-nums">
                          {formatBytes(tunnel.throughputOut)}
                        </div>
                        <div className="text-[10px] text-muted-foreground">
                          {formatTotalBytes(tunnel.cumulativeBytesOut)} total
                        </div>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
