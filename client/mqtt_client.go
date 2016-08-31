package client

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/evodicka/goccupied-client/io"
	"time"
)

var (
	config = io.ReadConfig()
	client = mqtt.NewClient(createClientOptions("goccupied"))
)

// Connects to the Adafruit MQTT Broker
func Connect() {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// Subscribes to the topic that contains the technical (0 and 1) occupied states and
// calls the passed handler function
func Subscribe(handler func(bool)) {

	f := func(client mqtt.Client, msg mqtt.Message) {
		free := bool(msg.Payload()[0]&1 == 1)
		handler(!free)
	}

	if token := client.Subscribe(config.Mqtt.ShareTopic, byte(0), f); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

// Sets the MQTT connections options
func createClientOptions(clientId string) *mqtt.ClientOptions {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Mqtt.Host)
	opts.SetUsername(config.Mqtt.User)
	opts.SetPassword(config.Mqtt.Password)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(10 * time.Second)
	opts.SetPingTimeout(10 * time.Second)

	return opts
}

// Disconnects from the Broker
func Disconnect() {
	client.Disconnect(250)
}
