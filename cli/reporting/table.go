package reporting

import (
	"green-go/lib"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func RenderTable(results []lib.HealthCheckResult) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Status", "Endpoint", "Protocol"})
	for _, result := range results {
		var healthSymbol string
		if result.Health {
			healthSymbol = "✅"
		} else {
			healthSymbol = "❌"
		}
		t.AppendRow([]interface{}{healthSymbol, result.Endpoint.Endpoint, result.Endpoint.Protocol})
	}

	t.Render()
}
