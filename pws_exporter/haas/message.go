package haas

import (
	"fmt"

	"github.com/AlekSi/pointer"
)

// https://www.home-assistant.io/integrations/sensor.mqtt/

// auto-discovery
// https://www.home-assistant.io/integrations/mqtt/#mqtt-discovery
// <discovery_prefix>/<component>/[<node_id>/]<object_id>/config
//{"unit_of_measurement":"%",
//"device_class":"humidity",
//"value_template":"{{ value_json.HUM }}",
//"state_class": "measurement",
//"state_topic":"rflink/Xiron-3201",
//"name":"eetkamer_humidity",
//"unique_id":"eetkamer_humidity",
//"device":{"identifiers":["xiron_3201"],
//"name":"xiron_3201",
//"model":"Digoo temp & humidity sensor","manufacturer":"Digoo"}}'

type Discovery struct {
	// origin
	// name
	// sw_version
	TopicPrefix       string  `json:"~"`
	Name              string  `json:"name"`
	StatT             string  `json:"stat_t"`
	UnitOfMeas        string  `json:"unit_of_meas"`
	DevCla            string  `json:"dev_cla"`
	FrcUpd            bool    `json:"frc_upd"`
	ValTpl            string  `json:"val_tpl"`
	Device            *Device `json:"device,omitempty"`
	AvailabilityTopic *string `json:"availability_topic,omitempty"`
	UniqueId          *string `json:"unique_id,omitempty"`

	ConfigTopic string
	// expire_after
	// retain
	// qos
	// availability_topic
	/*
		doc["name"] = "Plant " + String(sensorNumber) + " Humidity";
		doc["stat_t"]   = stateTopic;
		doc["unit_of_meas"] = "%";
		doc["dev_cla"] = "humidity";
		doc["frc_upd"] = true;
		doc["val_tpl"] = "{{ value_json.humidity|default(0) }}";
	*/
}

type Device struct {
	Name         string `json:"name"`
	Identifiers  string `json:"identifiers"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
}

func NewHumidityEntity(name string, uniqId string, topicPrefix string, device *Device) Discovery {
	return Discovery{
		topicPrefix,
		name,
		"~/state",
		"%",
		"humidity",
		true,
		"{{ value_json.humidity|default(0) }}",
		device,
		pointer.ToString("~/availability"),
		pointer.ToString(uniqId),
		fmt.Sprintf("%s/%s/config", topicPrefix, uniqId),
	}
}

func NewTemperatureEntity(name string, uniqId string, topicPrefix string, device *Device) Discovery {
	return Discovery{
		topicPrefix,
		name,
		"~/state",
		"°C",
		"temperature",
		true,
		"{{ value_json.temperature|default(0) }}",
		device,
		pointer.ToString("~/availability"),
		pointer.ToString(uniqId),
		fmt.Sprintf("%s/%s/config", topicPrefix, uniqId),
	}
}

type WeatherStation struct {
	IndoorTemperature  float32 `json:"indoor_temperature"`
	OutdoorTemperature float32 `json:"outdoor_temperature"`
	IndoorHumidity     uint    `json:"indoor_humidity"`
	OutdoorHumidity    uint    `json:"outdoor_humidity"`
	DewPoint           float32 `json:"dew_point"`
}

type TemperatureHumiditySensor struct {
	Temperature float32 `json:"temperature"`
	Humidity    uint    `json:"humidity"`
}
