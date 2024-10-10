package prometheus

import "github.com/prometheus/client_golang/prometheus"

var UpdatedGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_last_update",
		Help: "",
	}, []string{"station"},
)

var TemperatureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_outdoor_temperature",
		Help: "Outdoor temperature in Celsius",
	}, []string{"station"},
)

var IndoorTemperatureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_indoor_temperature",
		Help: "Indoor temperature in Celsius",
	}, []string{"station"},
)

var DewPointGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_dew_point",
		Help: "Dew point in Celsius",
	}, []string{"station"},
)

var HumidityGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_outdoor_humidity",
		Help: "Outdoor humidity in percentage",
	}, []string{"station"},
)

var IndoorHumidityGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_indoor_humidity",
		Help: "Indoor humidity in percentage",
	}, []string{"station"},
)

var WindDirGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_direction",
		Help: "Wind direction in degrees",
	}, []string{"station"},
)
var WindGustGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_gust",
		Help: "Wind gust in metres per second",
	}, []string{"station"},
)
var WindSpeedGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_wind_speed",
		Help: "Wind speed in metres per second",
	}, []string{"station"},
)

var RainGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_rain_hourly",
		Help: "Rain for last hour in millimeters",
	}, []string{"station"},
)
var RainDailyGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_rain_today",
		Help: "Rain for today in millimeters",
	}, []string{"station"},
)

var BarometricPressureGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_barometric_pressure",
		Help: "Barometric pressure in hecto Pascals",
	}, []string{"station"},
)
var UvIndexGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_uv_index",
		Help: "UV index",
	}, []string{"station"},
)
var SolarRadiationGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_solar_radiation",
		Help: "Solar radiation in watts per square metre",
	}, []string{"station"},
)

var SoilTempGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_soil_temperature",
		Help: "Soil temperature in Celsius",
	}, []string{"station", "channel"},
)

var SoilMoistGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "pws_soil_moisture",
		Help: "Soil moisture in percentage",
	}, []string{"station", "channel"},
)

func RegisterMetrics() {
	prometheus.MustRegister(UpdatedGauge)
	prometheus.MustRegister(TemperatureGauge)
	prometheus.MustRegister(IndoorTemperatureGauge)
	prometheus.MustRegister(DewPointGauge)
	prometheus.MustRegister(HumidityGauge)
	prometheus.MustRegister(IndoorHumidityGauge)
	prometheus.MustRegister(WindDirGauge)
	prometheus.MustRegister(WindGustGauge)
	prometheus.MustRegister(WindSpeedGauge)
	prometheus.MustRegister(RainGauge)
	prometheus.MustRegister(RainDailyGauge)
	prometheus.MustRegister(BarometricPressureGauge)
	prometheus.MustRegister(UvIndexGauge)
	prometheus.MustRegister(SolarRadiationGauge)
	prometheus.MustRegister(SoilTempGauge)
	prometheus.MustRegister(SoilMoistGauge)
}
