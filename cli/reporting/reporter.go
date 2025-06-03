package reporting

import (
	"encoding/json"
	"fmt"
	"green-go/lib"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Reporter interface {
	Render(results []lib.HealthCheckResult)
}

type Table struct {
}

type Json struct {
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

type HealthCheckResultJson struct {
	Status   bool   `json:"status"`
	Endpoint string `json:"endpoint"`
	Protocol string `json:"string"`
}

func (reporter *Json) Render(results []lib.HealthCheckResult) {
	datas := make([]HealthCheckResultJson, len(results))
	for i, result := range results {
		datas[i] = HealthCheckResultJson{
			Status:   result.Health,
			Endpoint: result.Endpoint.Endpoint,
			Protocol: result.Endpoint.Protocol,
		}
	}

	jsonByte, err := json.Marshal(datas)
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonByte)

	fmt.Println(jsonString)
}

func GetByType(reporterType string) Reporter {
	switch reporterType {
	case "table":
		{
			return &Table{}
		}
	case "json":
		{
			return &Json{}
		}
	}

	return nil
}
