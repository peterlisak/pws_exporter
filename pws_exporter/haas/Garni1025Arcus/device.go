package Garni1025Arcus

import (
	"fmt"

	"github.com/peterlisak/pws_exporter/pws_exporter/haas"
)

type Device struct {
	device             haas.Device
	IndoorTemperature  haas.Discovery
	IndoorHumidity     haas.Discovery
	OutdoorTemperature haas.Discovery
	OutdoorHumidity    haas.Discovery
	DewPoint           haas.Discovery
}

func New(id int) *Device {
	name := fmt.Sprintf("pws-%d", id)
	device := &haas.Device{
		name,
		fmt.Sprintf("Garni-1025-Arcus_%s", name),
		"1025 Arcus",
		"Garni",
	}
	topicPrefix := fmt.Sprintf("homeassistant/sensor/%s", device.Name)
	return &Device{
		*device,
		haas.NewTemperatureEntity(fmt.Sprintf("in-temp-%d", id), fmt.Sprintf("in-temp-%d_%d", id, 1), topicPrefix, device),
		haas.NewHumidityEntity(fmt.Sprintf("in-hum-%d", id), fmt.Sprintf("in-hum-%d_%d", id, 1), topicPrefix, device),
		haas.NewTemperatureEntity(fmt.Sprintf("out-temp-%d", id), fmt.Sprintf("out-temp-%d_%d", id, 2), topicPrefix, device),
		haas.NewHumidityEntity(fmt.Sprintf("out-hum-%d", id), fmt.Sprintf("hum-%d_%d", id, 2), topicPrefix, device),
		haas.NewTemperatureEntity(fmt.Sprintf("dew-point-%d", id), fmt.Sprintf("dew-point-%d_%d", id, 1), topicPrefix, device),
	}
}

func (d *Device) Entities() []haas.Discovery {
	return []haas.Discovery{
		d.IndoorTemperature,
		d.IndoorHumidity,
		d.OutdoorTemperature,
		d.OutdoorHumidity,
		d.DewPoint,
	}
}

func (d *Device) StateTopic() string {
	return fmt.Sprintf("homeassistant/sensor/%s/state", d.device.Name)
}
