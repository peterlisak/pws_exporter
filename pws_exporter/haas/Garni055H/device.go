package Garni055H

import (
	"fmt"

	"github.com/peterlisak/pws_exporter/pws_exporter/haas"
)

type Device struct {
	device      haas.Device
	Temperature haas.Discovery
	Humidity    haas.Discovery
}

func New(id int) *Device {
	name := fmt.Sprintf("temp-hum-sensor-%d", id)
	device := &haas.Device{
		name,
		fmt.Sprintf("Garni-055H_%s", name),
		"055H",
		"Garni",
	}
	topicPrefix := fmt.Sprintf("homeassistant/sensor/%s", device.Name)
	return &Device{
		*device,
		haas.NewTemperatureEntity("temp", fmt.Sprintf("temp-%d_%d", id, 1), topicPrefix, device),
		haas.NewHumidityEntity("hum", fmt.Sprintf("hum-%d_%d", id, 2), topicPrefix, device),
	}
}

func (d *Device) Entities() []haas.Discovery {
	return []haas.Discovery{
		d.Temperature,
		d.Humidity,
	}
}

func (d *Device) StateTopic() string {
	return fmt.Sprintf("homeassistant/sensor/%s/state", d.device.Name)
}
