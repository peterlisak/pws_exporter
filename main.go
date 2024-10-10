package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterlisak/pws_exporter/pws_exporter/haas"
	"github.com/peterlisak/pws_exporter/pws_exporter/haas/Garni055H"
	"github.com/peterlisak/pws_exporter/pws_exporter/haas/Garni1025Arcus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/peterlisak/pws_exporter/pws_exporter"
	"github.com/peterlisak/pws_exporter/pws_exporter/prometheus"
)

func main() {
	// MQTT
	client := pws_exporter.InitMQTTClient()
	defer client.Disconnect(250)
	fmt.Println("MQTT client ready")
	// Prometheus
	prometheus.RegisterMetrics()
	// HAaS
	pws := haas.PwsDevice{
		Garni1025Arcus.New(110),
		[]haas.HaasDevice{
			Garni055H.New(111),
			Garni055H.New(112),
			Garni055H.New(113),
		},
	}
	// Discovery
	pws_exporter.Discover(client, pws)
	// Server
	handler := pws_exporter.NewHandler(client, pws)
	http.HandleFunc("/", handler.HandleHome)
	http.HandleFunc("/weatherstation/updateweatherstation.php", handler.HandleUpdatePWS)
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
