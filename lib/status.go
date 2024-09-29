package lib

import (
	"fmt"
	"log"
	"net/http"
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

func (e Endpoint) Check(health chan bool) {

	switch e.Protocol {
	case "http":
		{
			resp, err := http.Get("http://" + e.Endpoint + ":" + fmt.Sprint(e.Port))
			if err != nil {
				log.Println("Error getting " + e.Endpoint)
				health <- false
			}
			health <- resp.StatusCode == e.Status
		}
	case "https":
		{
			resp, err := http.Get("https://" + e.Endpoint + ":" + fmt.Sprint(e.Port))
			if err != nil {
				log.Println("Error getting " + e.Endpoint)
				health <- false
			}
			health <- resp.StatusCode == e.Status
		}
	case "ping":
		log.Println("Ping is not implemented yet.")
		health <- false
	}
	health <- false
}
