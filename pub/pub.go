package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type PubOpt struct {
	topic     string // The topic name to/from which to publish/subscribe
	broker    string // The broker URI. ex: tcp://10.10.1.1:1883
	password  string // The password (optional)
	user      string // The User (optional)
	id        string // The ClientID (optional)
	cleansess bool   // Set Clean Session (default false)
	qos       int    // The Quality of Service 0,1,2 (default 0)
	num       int    // The number of messages to publish or subscribe (default 1)
	payload   string // The message text to publish (default empty)
}

func main() {
	pubOpt := PubOpt{
		topic:     "test",
		broker:    "tcp://192.168.64.255:1883",
		password:  "",
		user:      "",
		id:        "pub1",
		cleansess: false,
		qos:       0,
		num:       1,
	}
	fmt.Printf("Publisher Info:\n")
	fmt.Printf("\tbroker:    %s\n", pubOpt.broker)
	fmt.Printf("\tclientid:  %s\n", pubOpt.id)
	fmt.Printf("\tuser:      %s\n", pubOpt.user)
	fmt.Printf("\tpassword:  %s\n", pubOpt.password)
	fmt.Printf("\ttopic:     %s\n", pubOpt.topic)
	fmt.Printf("\tqos:       %d\n", pubOpt.qos)
	fmt.Printf("\tcleansess: %v\n", pubOpt.cleansess)
	fmt.Printf("\tnum:       %d\n", pubOpt.num)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(pubOpt.broker)
	opts.SetClientID(pubOpt.id)
	opts.SetUsername(pubOpt.user)
	opts.SetPassword(pubOpt.password)
	opts.SetCleanSession(pubOpt.cleansess)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Publisher started")
	var payload string

	for {
		fmt.Scan(&payload)
		pubOpt.payload = payload
		for i := 0; i < pubOpt.num; i++ {
			fmt.Println("---- doing publish ----")
			token := client.Publish(pubOpt.topic, byte(pubOpt.qos), false, pubOpt.payload)
			token.Wait()
		}
	}
}
