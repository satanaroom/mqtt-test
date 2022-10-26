package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubOpt struct {
	topic     string // The topic name to/from which to publish/subscribe
	broker    string // The broker URI. ex: tcp://10.10.1.1:1883
	password  string // The password (optional)
	user      string // The User (optional)
	id        string // The ClientID (optional)
	cleansess bool   // Set Clean Session (default false)
	qos       int    // The Quality of Service 0,1,2 (default 0)
	num       int    // The number of messages to publish or subscribe (default 1)
	store     string // The Store Directory (default use memory store)
}

func main() {
	subOpt := SubOpt{
		topic:     "test",
		broker:    "tcp://localhost:1883",
		password:  "qwerty1",
		user:      "ftragula",
		id:        "sub1",
		cleansess: false,
		qos:       0,
		num:       1,
		store:     "/home/srnpo/GolandProjects/mqtt-test/store",
	}
	fmt.Printf("Subscriber Info:\n")
	fmt.Printf("\tbroker:    %s\n", subOpt.broker)
	fmt.Printf("\tclientid:  %s\n", subOpt.id)
	fmt.Printf("\tuser:      %s\n", subOpt.user)
	fmt.Printf("\tpassword:  %s\n", subOpt.password)
	fmt.Printf("\ttopic:     %s\n", subOpt.topic)
	fmt.Printf("\tqos:       %d\n", subOpt.qos)
	fmt.Printf("\tcleansess: %v\n", subOpt.cleansess)
	fmt.Printf("\tnum:       %d\n", subOpt.num)
	fmt.Printf("\tstore:     %s\n", subOpt.store)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(subOpt.broker)
	opts.SetClientID(subOpt.id)
	opts.SetUsername(subOpt.user)
	opts.SetPassword(subOpt.password)
	opts.SetCleanSession(subOpt.cleansess)
	if subOpt.store != ":memory:" {
		opts.SetStore(mqtt.NewFileStore(subOpt.store))
	}

	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(subOpt.topic, byte(subOpt.qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Println("Subscriber started")

	for {
		incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
	}
}
