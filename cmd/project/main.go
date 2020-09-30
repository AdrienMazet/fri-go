package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fri-go/internal/paho"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func redisSubHandler(client mqtt.Client, message mqtt.Message) {
	// install redis
	// configure redis (port, credentials ?, persistence)
	// tests ?
	redis.StoreData(message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	client := paho.Connect("tcp://localhost:1883", "mqtt_client")
	topic := "main"
	if token := client.Subscribe(topic, 0, redisSubHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	paho.StartSensorsPubs(client, topic, 10)

	<-c
}
