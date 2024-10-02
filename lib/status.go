package lib

import (
	"log"
	"net/http"
	"os/exec"
)

type Endpoint struct {
	Endpoint string `yaml:"endpoint"`
	Protocol string `yaml:"protocol"`
	Port     int    `yaml:"port"`
	Status   int    `yaml:"status"`
}

type HealthCheckResult struct {
	Endpoint Endpoint
	Health   bool
}

func PerformChecks(endpoints []Endpoint) []HealthCheckResult {
	results := make([]HealthCheckResult, len(endpoints))

	for i := range len(endpoints) {

		health := make(chan bool)
		go endpoints[i].Check(health)

		result := HealthCheckResult{}
		result.Endpoint = endpoints[i]

		result.Health = <-health

		results[i] = result
	}

	return results
}

func checkHttp(e Endpoint, health chan bool) {
	resp, err := http.Get(e.Endpoint)
	if err != nil {
		log.Println("Error getting " + e.Endpoint)
		health <- false
		return
	}
	_ = resp.Body.Close()
	health <- resp.StatusCode == e.Status
}

func checkPing(e Endpoint, health chan bool) {
	cmd := exec.Command("ping", "-c", "1", "-W", "2", e.Endpoint)
	health <- cmd.Run() == nil
}

func (e Endpoint) Check(health chan bool) {

	switch e.Protocol {
	case "http":
		{
			checkHttp(e, health)
		}
	case "https":
		{
			checkHttp(e, health)
		}
	case "ping":
		checkPing(e, health)
	}
	health <- false
}
