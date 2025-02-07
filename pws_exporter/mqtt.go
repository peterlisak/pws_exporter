package pws_exporter

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peterlisak/pws_exporter/pws_exporter/haas"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	/*
		if string(msg.Payload()) == "online" {
			// Discovery
			Discover(client)
		}
	*/
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func InitMQTTClient() mqtt.Client {
	var broker = "peter-pn40"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("admin")
	opts.SetPassword("admin")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("panic %v", token.Error()))
	}

	return client
}

func Publish(client mqtt.Client, topic, msg string, retained bool) {
	token := client.Publish(topic, 0, retained, msg)
	// todo WaitTimeout
	token.Wait()
}

func subscribe(client mqtt.Client) {
	// subscribe to the same topic, that was published to, to receive the messages
	topic := "homeassistant/status"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	// Check for errors during subscribe (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		fmt.Printf("Failed to subscribe to topic")
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

func Discover(client mqtt.Client, pws haas.PwsDevice) {
	publishDeviceEntities(client, pws.Pws.Entities())
	for _, sensor := range pws.Sensors {
		publishDeviceEntities(client, sensor.Entities())
	}
}

func publishDeviceEntities(client mqtt.Client, entities []haas.Discovery) {
	var data []byte
	for _, ent := range entities {
		data, _ = json.Marshal(ent)
		fmt.Println(ent.ConfigTopic, string(data))
		Publish(client, ent.ConfigTopic, string(data), true)
	}
}
