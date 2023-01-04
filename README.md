# pws_exporter

This is GO server for getting weather information from your Personal Weather Station. It exposes the latest 
received weather information as metrics for Prometheus scrapper.

Note: All units are converted to metrics units.

Exposed endpoints:
- `/`: help (TODO)
- `/weatherstation/updateweatherstation.php`: for PWS to send weather data
- `/metrics`: exposing Prometheus-like metrics for Prometheus scrapper

## Supported parameters from PWS

Currently, these parametres based on [GARNI 1025 ARCUS](https://www.garnitechnology.cz/garni-1025-arcus/) are supported.
Based on [PWS Upload API](https://support.weather.com/s/article/PWS-Upload-Protocol?language=en_US).

- `ID`: ID as registered by wunderground.com 
  - used as `station` label for all 
- `PASSWORD`: Station Key registered with this PWS ID by wunderground.com (ignored at the moment)
- `dateutc`: In YYYY-MM-DD HH:MM:SS format and UTC time zone
- `winddir`: 0-360 instantaneous wind direction 
  - exposed as `pws_wind_direction`
- `windspeedmph`: mph instantaneous wind speed (converted to metres per second) 
  - exposed as `pws_wind_speed`
- `windgustmph`: mph current wind gust, using software specific time period (converted to metres per second) 
  - exposed as `pws_wind_gust`
- `humidity`: outdoor humidity in range 0-100% 
  - exposed as `pws_humidity`
- `dewptf`: outdoor dewpoint in Fahrenheit (converted to Celsius) 
  - exposed as `pws_dew_point`
- `tempf`: outdoor temperature in Fahrenheit (converted to Celsius) 
  - exposed as `pws_temperature`
- `rainin`: accumulated rainfall inches over the past hour (converted to millimetres) 
  - exposed as `pws_rain_hourly`
- `dailyrainin`: accumulated rainfall inches so far today in local time (converted to millimetres) 
  - exposed as `pws_rain_today`
- `baromin`: barometric pressure inches (converted to hecto Pascals) 
  - exposed as `pws_barometric_pressure`
- `soiltempf`: soil temperature in Fahrenheit (converted to Celsius) 
  - for extra sensors use soiltemp2f, soiltemp3f, and soiltemp4f
  - exposed as `pws_soil_temperature`, extra sensors are distinguished by label `sensor`
- `soilmoisture`: soil moisture in range 0-100%
  - for extra sensors use soilmoisture2, soilmoisture3, and soilmoisture4
  - exposed as `pws_soil_moisture`, extra sensors are distinguished by label `sensor`
- `solarradiation`: solar radiation in Watts per square meter
  - exposed as `pws_solar_radiation`
- `UV`: UV index 
  - exposed as `pws_uv_index`
- `indoortempf`: indoor temperature in Fahrenheit (converted to Celsius) 
  - exposed as `pws_indoor_temperature`
- `indoorhumidity`: indoor humidity in range 0-100% 
  - exposed as `pws_indoor_humidity`

NOTE: `soiltempf` and `soilmoisture` fields are used for data from other sensors (not only soil sensors).

## TODO
- basic configuration (e.g. port, etc.)
- Help on `/`

## Release new version

```
VERSION=vX.Y.Z
git commit -m "mymodule: changes for $VERSION"
git tag $VERSION

git push origin $VERSION

GOPROXY=proxy.golang.org go list -m github.com/peterlisak/pws_exporter@$VERSION
```
