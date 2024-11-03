package reporting

import (
	"green-go/lib"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Reporter interface {
	Render(results []lib.HealthCheckResult)
}

type Table struct {
}

func (reporter *Table) Render(results []lib.HealthCheckResult) {

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

func GetByType(reporterType string) Reporter {
	if reporterType == "table" {
		return &Table{}
	}

	return nil
}
