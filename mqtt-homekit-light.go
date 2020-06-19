package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	var (
		accessoryManu = flag.String("manufacturer", "hc", "homekit accessory manufacturer")
		accessoryName = flag.String("name", "Light", "homekit accessory name")
		mqttClientID  = flag.String("mqtt-client-id", "homekit-mqtt", "mqtt client ID")
		mqttTopic     = flag.String("mqtt-topic", "homekit/light", "topic to listen for state, will call <mqtt-topic>/set when homekit triggers")
		mqttURL       = flag.String("mqtt", "mqtt://localhost:1883", "mqtt url")
		pin           = flag.String("pin", "32191123", "homekit PIN for pairing")
		storagePath   = flag.String("storage-path", "./", "where to store persistent files")
	)
	flag.Parse()

	co := mqtt.NewClientOptions()
	co.AddBroker(*mqttURL)
	co.SetClientID(*mqttClientID)
	mqcli := mqtt.NewClient(co)
	if t := mqcli.Connect(); t.Wait() && t.Error() != nil {
		logrus.Fatalf("Failed to connect to mqtt: %q", t.Error())
	}

	acc := accessory.NewLightbulb(accessory.Info{
		Name:         *accessoryName,
		Manufacturer: *accessoryManu,
	})

	var lightHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
		topic := m.Topic()
		message, err := strconv.ParseBool(string(m.Payload()))
		if err != nil {
			logrus.Errorf("Unable to parse message as boolean: %q", err)
		}

		acc.Lightbulb.On.SetValue(message)
		logrus.Debugf("mqtt: %s %v", topic, message)
	}

	if t := mqcli.Subscribe(*mqttTopic, 0, lightHandler); t.Wait() && t.Error() != nil {
		logrus.Errorf("Error: subscribing %q", t.Error())
	}

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if t := mqcli.Publish(fmt.Sprintf("%s/set", *mqttTopic), 0, false, fmt.Sprintf("%v", on)); t.Wait() && t.Error() != nil {
			logrus.Errorf("Error: publishing %q", t.Error())
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: *pin, StoragePath: *storagePath}, acc.Accessory)
	if err != nil {
		logrus.Fatal(err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
