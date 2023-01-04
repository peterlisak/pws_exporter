package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math"
	"net/http"
)

type Params struct {
	// mandatory
	Id string `schema:"ID,required"`
	/* Password string `schema:"PASSWORD,required"`
	Dateutc  string `schema:"dateutc,required"` */
	// optionals
	/* Action         string
	Realtime       bool */
	Rtfreq         float32
	Baromin        float32
	Tempf          float32
	Dewptf         float32
	Humidity       int
	Windspeedmph   float32
	Windgustmph    float32
	Winddir        int
	Rainin         float32
	Dailyrainin    float32
	Solarradiation float32
	UV             float32
	Indoortempf    float32
	Indoorhumidity int
	Soiltempf      float32
	Soilmoisture   int
	Soiltemp2f     float32
	Soilmoisture2  int
	Soiltemp3f     float32
	Soilmoisture3  int
}

var updatedGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_last_update",
		Help: "",
	}, []string{"station"},
)

var temperatureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_outdoor_temperature",
		Help: "Outdoor temperature in Celsius",
	}, []string{"station"},
)

var indoorTemperatureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_indoor_temperature",
		Help: "Indoor temperature in Celsius",
	}, []string{"station"},
)

var dewPointGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_dew_point",
		Help: "Dew point in Celsius",
	}, []string{"station"},
)

var humidityGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_outdoor_humidity",
		Help: "Outdoor humidity in percentage",
	}, []string{"station"},
)

var indoorHumidityGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_indoor_humidity",
		Help: "Indoor humidity in percentage",
	}, []string{"station"},
)

var windDirGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_direction",
		Help: "Wind direction in degrees",
	}, []string{"station"},
)
var windGustGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_gust",
		Help: "Wind gust in metres per second",
	}, []string{"station"},
)
var windSpeedGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_speed",
		Help: "Wind speed in metres per second",
	}, []string{"station"},
)

var rainGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_rain_hourly",
		Help: "Rain for last hour in millimeters",
	}, []string{"station"},
)
var rainDailyGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_rain_today",
		Help: "Rain for today in millimeters",
	}, []string{"station"},
)

var barometricPressureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_barometric_pressure",
		Help: "Barometric pressure in hecto Pascals",
	}, []string{"station"},
)
var uvIndexGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_uv_index",
		Help: "UV index",
	}, []string{"station"},
)
var solarRadiationGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_solar_radiation",
		Help: "Solar radiation in watts per square metre",
	}, []string{"station"},
)

var soilTempGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_soil_temperature",
		Help: "Soil temperature in Celsius",
	}, []string{"station", "channel"},
)

var soilMoistGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_soil_moisture",
		Help: "Soil moisture in percentage",
	}, []string{"station", "channel"},
)

var decoder = schema.NewDecoder()

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PWS data to Promethes\n")
}

func handleParams(w http.ResponseWriter, r *http.Request) {
	var params = Params{}

	// r.PostForm is a map of our POST form values
	decoder.IgnoreUnknownKeys(true)
	query := r.URL.Query()
	var err = decoder.Decode(&params, query)
	if err != nil {
		// Handle error
		log.Printf("Missing %s", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad request")
		return
	}

	log.Printf("Params: %v", r.URL.Query())
	//fmt.Fprintf(w, "Params: %v", r.URL.Query())
	updatedGauge.WithLabelValues(params.Id).SetToCurrentTime()
	temperatureGauge.WithLabelValues(params.Id).Set(float64((params.Tempf - 32) / 1.8))
	indoorTemperatureGauge.WithLabelValues(params.Id).Set(float64((params.Indoortempf - 32) / 1.8))
	dewPointGauge.WithLabelValues(params.Id).Set(float64((params.Dewptf - 32) / 1.8))
	humidityGauge.WithLabelValues(params.Id).Set(float64(params.Humidity))
	indoorHumidityGauge.WithLabelValues(params.Id).Set(float64(params.Indoorhumidity))
	windDirGauge.WithLabelValues(params.Id).Set(float64(params.Winddir))
	windGustGauge.WithLabelValues(params.Id).Set(float64(params.Windgustmph * 1.60934 / 3.6))
	windSpeedGauge.WithLabelValues(params.Id).Set(float64(params.Windspeedmph * 1.60934 / 3.6))
	rainGauge.WithLabelValues(params.Id).Set(float64(params.Rainin * 25.4))
	rainDailyGauge.WithLabelValues(params.Id).Set(float64(params.Dailyrainin * 25.4))
	barometricPressureGauge.WithLabelValues(params.Id).Set(math.Round(float64(params.Baromin) * 33.863886666667))
	uvIndexGauge.WithLabelValues(params.Id).Set(float64(params.UV))
	solarRadiationGauge.WithLabelValues(params.Id).Set(float64(params.Solarradiation))

	soilMoistGauge.Reset()
	if query.Has("soilmoisture") {
		soilMoistGauge.WithLabelValues(params.Id, "1").Set(float64(params.Soilmoisture))
	}
	if query.Has("soilmoisture2") {
		soilMoistGauge.WithLabelValues(params.Id, "2").Set(float64(params.Soilmoisture2))
	}
	if query.Has("soilmoisture3") {
		soilMoistGauge.WithLabelValues(params.Id, "3").Set(float64(params.Soilmoisture3))
	}

	soilTempGauge.Reset()
	if query.Has("soiltempf") {
		soilTempGauge.WithLabelValues(params.Id, "1").Set(float64((params.Soiltempf - 32) / 1.8))
	}
	if query.Has("soiltemp2f") {
		soilTempGauge.WithLabelValues(params.Id, "2").Set(float64((params.Soiltemp2f - 32) / 1.8))
	}
	if query.Has("soiltemp3f") {
		soilTempGauge.WithLabelValues(params.Id, "3").Set(float64((params.Soiltemp3f - 32) / 1.8))
	}
	fmt.Fprint(w, "success")
}

func main() {
	prometheus.MustRegister(updatedGauge)
	prometheus.MustRegister(temperatureGauge)
	prometheus.MustRegister(indoorTemperatureGauge)
	prometheus.MustRegister(dewPointGauge)
	prometheus.MustRegister(humidityGauge)
	prometheus.MustRegister(indoorHumidityGauge)
	prometheus.MustRegister(windDirGauge)
	prometheus.MustRegister(windGustGauge)
	prometheus.MustRegister(windSpeedGauge)
	prometheus.MustRegister(rainGauge)
	prometheus.MustRegister(rainDailyGauge)
	prometheus.MustRegister(barometricPressureGauge)
	prometheus.MustRegister(uvIndexGauge)
	prometheus.MustRegister(solarRadiationGauge)
	prometheus.MustRegister(soilTempGauge)
	prometheus.MustRegister(soilMoistGauge)
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/weatherstation/updateweatherstation.php", handleParams)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
