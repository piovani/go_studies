package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/piovani/go_api/prometheus/pkg/metric"
)

func handleParams() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Invalid query")
	}

	return os.Args[1], nil
}

func main() {
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	appMetric := metric.NewCLI("search")
	appMetric.Started()

	query, err := handleParams()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(query)

	appMetric.Finished()
	if err = metricService.SaveCLI(appMetric); err != nil {
		log.Fatal(err.Error())
	}
}
