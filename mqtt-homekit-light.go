package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peterbourgon/ff/v3"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	var (
		fs    = flag.NewFlagSet("mqtt-homekit-light", flag.ExitOnError)
		_     = fs.String("config", "", "Config file to load settings from")
		debug = fs.Bool("debug", false, "Enable debugging")

		accessoryManu = fs.String("manufacturer", "hc", "homekit accessory manufacturer")
		accessoryName = fs.String("name", "Light", "homekit accessory name")
		mqttClientID  = fs.String("mqtt-client-id", "homekit-mqtt", "mqtt client ID")
		mqttTopic     = fs.String("mqtt-topic", "homekit/light", "topic to listen for state, will call <mqtt-topic>/set when homekit triggers")
		mqttURL       = fs.String("mqtt", "mqtt://localhost:1883", "mqtt url")
		pin           = fs.String("pin", "32191123", "homekit PIN for pairing")
		storagePath   = fs.String("storage-path", "./", "where to store persistent files")
	)
	logrus.Info("config")
	ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix("MHL"), ff.WithConfigFileFlag("config"), ff.WithConfigFileParser(ff.JSONParser))

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

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
		logrus.Debugf("hap: %v", on)

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
