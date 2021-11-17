package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Telemetry struct {
	AgentID       string                 `json:"agent_id"`
	Version       string                 `json:"version"`
	Payload       map[string]interface{} `json:"payload"`
	CloudProvider string                 `json:"cloud_provider"`
	Timestamp     time.Time              `json:"timestamp"`
}

func collectHandler(client influxdb2.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var telemetry Telemetry

		err := json.NewDecoder(r.Body).Decode(&telemetry)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		payload := telemetry.Payload
		fmt.Println(telemetry)
		writeAPI := client.WriteAPIBlocking("trento", "telemetry")
		p := influxdb2.NewPoint("telemetry",
			map[string]string{"app": "trento", "version": telemetry.Version, "agent_id": telemetry.AgentID, "cloud_provider": telemetry.CloudProvider},
			payload,
			telemetry.Timestamp)

		err = writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func handleRequests(client influxdb2.Client) {
	http.HandleFunc("/api/collect", collectHandler(client))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	client := influxdb2.NewClient(os.Getenv("TELEMETRY_INFLUXDB_URL"), os.Getenv("TELEMETRY_INFLUXDB_TOKEN"))
	defer client.Close()

	handleRequests(client)
}
