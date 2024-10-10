package pws_exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/schema"
	"github.com/peterlisak/pws_exporter/pws_exporter/haas"
	"github.com/peterlisak/pws_exporter/pws_exporter/prometheus"
)

type Handler struct {
	Decoder    *schema.Decoder
	MQTTClient mqtt.Client
	Pws        haas.PwsDevice
}

func NewHandler(client mqtt.Client, pws haas.PwsDevice) *Handler {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	return &Handler{
		decoder,
		client,
		pws,
	}
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PWS data to Promethes\n")
}

func (h *Handler) HandleUpdatePWS(w http.ResponseWriter, r *http.Request) {
	var params = PWSParams{}

	// r.PostForm is a map of our POST form values
	//decoder.IgnoreUnknownKeys(true)
	query := r.URL.Query()
	var err = h.Decoder.Decode(&params, query)
	if err != nil {
		// Handle error
		log.Printf("Missing %s", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	log.Printf("Params: %v", r.URL.Query())
	prometheus.UpdatedGauge.WithLabelValues(params.Id).SetToCurrentTime()
	prometheus.TemperatureGauge.WithLabelValues(params.Id).Set(float64((params.Tempf - 32) / 1.8))
	prometheus.IndoorTemperatureGauge.WithLabelValues(params.Id).Set(float64((params.Indoortempf - 32) / 1.8))
	prometheus.DewPointGauge.WithLabelValues(params.Id).Set(float64((params.Dewptf - 32) / 1.8))
	prometheus.HumidityGauge.WithLabelValues(params.Id).Set(float64(params.Humidity))
	prometheus.IndoorHumidityGauge.WithLabelValues(params.Id).Set(float64(params.Indoorhumidity))
	prometheus.WindDirGauge.WithLabelValues(params.Id).Set(float64(params.Winddir))
	prometheus.WindGustGauge.WithLabelValues(params.Id).Set(float64(params.Windgustmph * 1.60934 / 3.6))
	prometheus.WindSpeedGauge.WithLabelValues(params.Id).Set(float64(params.Windspeedmph * 1.60934 / 3.6))
	prometheus.RainGauge.WithLabelValues(params.Id).Set(float64(params.Rainin * 25.4))
	prometheus.RainDailyGauge.WithLabelValues(params.Id).Set(float64(params.Dailyrainin * 25.4))
	prometheus.BarometricPressureGauge.WithLabelValues(params.Id).Set(math.Round(float64(params.Baromin) * 33.863886666667))
	prometheus.UvIndexGauge.WithLabelValues(params.Id).Set(float64(params.UV))
	prometheus.SolarRadiationGauge.WithLabelValues(params.Id).Set(float64(params.Solarradiation))

	prometheus.SoilMoistGauge.Reset()
	if query.Has("soilmoisture") {
		prometheus.SoilMoistGauge.WithLabelValues(params.Id, "1").Set(float64(params.Soilmoisture))
	}
	if query.Has("soilmoisture2") {
		prometheus.SoilMoistGauge.WithLabelValues(params.Id, "2").Set(float64(params.Soilmoisture2))
	}
	if query.Has("soilmoisture3") {
		prometheus.SoilMoistGauge.WithLabelValues(params.Id, "3").Set(float64(params.Soilmoisture3))
	}

	prometheus.SoilTempGauge.Reset()
	if query.Has("soiltempf") {
		prometheus.SoilTempGauge.WithLabelValues(params.Id, "1").Set(float64((params.Soiltempf - 32) / 1.8))
	}
	if query.Has("soiltemp2f") {
		prometheus.SoilTempGauge.WithLabelValues(params.Id, "2").Set(float64((params.Soiltemp2f - 32) / 1.8))
	}
	if query.Has("soiltemp3f") {
		prometheus.SoilTempGauge.WithLabelValues(params.Id, "3").Set(float64((params.Soiltemp3f - 32) / 1.8))
	}
	fmt.Fprint(w, "success")

	//subscribe(h.MQTTClient)
	msg, _ := json.Marshal(haas.WeatherStation{
		(params.Indoortempf - 32) / 1.8,
		(params.Tempf - 32) / 1.8,
		uint(params.Indoorhumidity),
		uint(params.Humidity),
		(params.Dewptf - 32) / 1.8,
	})
	Publish(h.MQTTClient, h.Pws.Pws.StateTopic(), string(msg), false)

	if query.Has("soilmoisture") && query.Has("soiltempf") {
		msg, _ := json.Marshal(haas.TemperatureHumiditySensor{
			(params.Soiltempf - 32) / 1.8,
			uint(params.Soilmoisture),
		})
		Publish(h.MQTTClient, h.Pws.Sensors[0].StateTopic(), string(msg), false)
	}
	if query.Has("soilmoisture2") && query.Has("soiltempf2") {
		msg, _ := json.Marshal(haas.TemperatureHumiditySensor{
			(params.Soiltemp2f - 32) / 1.8,
			uint(params.Soilmoisture2),
		})
		Publish(h.MQTTClient, h.Pws.Sensors[1].StateTopic(), string(msg), false)
	}
	if query.Has("soilmoisture3") && query.Has("soiltempf3") {
		msg, _ := json.Marshal(haas.TemperatureHumiditySensor{
			(params.Soiltemp3f - 32) / 1.8,
			uint(params.Soilmoisture3),
		})
		Publish(h.MQTTClient, h.Pws.Sensors[2].StateTopic(), string(msg), false)
	}
}
