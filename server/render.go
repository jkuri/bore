package server

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/jkuri/bore/internal/version"
)

const (
	black      = lipgloss.Color("0")
	lightGreen = lipgloss.Color("10")
	lightBlue  = lipgloss.Color("12")
	gray       = lipgloss.Color("245")
	lightGray  = lipgloss.Color("241")
)

func renderTable(data clientResponse, w io.Writer) {
	re := lipgloss.NewRenderer(w)

	var (
		HeaderStyle  = re.NewStyle().Foreground(lightBlue).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1).Width(20).Align(lipgloss.Center)
		OddRowStyle  = CellStyle.Copy().Foreground(gray)
		EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
		BorderStyle  = lipgloss.NewStyle().Foreground(lightBlue)
	)

	rows := [][]string{
		{"HTTP", fmt.Sprintf("http://%s.%s", data.id, data.domain)},
		{"HTTPS", fmt.Sprintf("https://%s.%s", data.id, data.domain)},
		{"TCP", fmt.Sprintf("tcp://%s:%d", data.domain, data.port)},
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			// Make the second column a little wider.
			if col == 1 {
				style = style.Copy().Width(60)
			}

			return style
		}).
		Headers("Protocol", "URL").
		Rows(rows...)

	io.WriteString(w, t.String())
	io.WriteString(w, "\n")
}

func renderMessage(data clientResponse, w io.Writer) {
	style := lipgloss.NewStyle().Bold(true).Foreground(lightGreen).Background(black)
	io.WriteString(w, style.Render("Welcome to bore server", version.Version, "at", data.domain))
	io.WriteString(w, "\n")
}
