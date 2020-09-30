package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"../../internal/paho"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	// process message here
	// redis
	// CSV
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	client := paho.Connect("tcp://localhost:1883", "mqtt_client")
	topic := "main"
	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	paho.StartSensorsPubs(client, topic, 10)

	<-c
}
