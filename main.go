package main

import (
	"fmt"
	"net/url"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient() (mqtt.Client, error) {
	mqttUrl := url.URL{
		Scheme: "upd",
		Host:   "192.168.64.255:1883",
	}
	clientOpts := mqtt.ClientOptions{
		Servers: []*url.URL{&mqttUrl},
	}
	client := mqtt.NewClient(&clientOpts)

	return client, nil
}

func main() {
	mqttClient, err := NewMqttClient()
	if err != nil {
		panic(err)
	}
	token := mqttClient.Connect()
	fmt.Println(token)
}
