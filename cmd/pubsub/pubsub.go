package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fri-go/internal/configuration"
	"github.com/fri-go/internal/csv"
	"github.com/fri-go/internal/paho"
	"github.com/fri-go/internal/redis"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	go redis.StoreData(message.Payload())
	go csv.StoreData(message.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conf := configuration.LoadConfiguration()

	client := paho.Connect(conf.Adress+":"+conf.Port, conf.ClientID)

	for i := 0; i < len(conf.Airports); i++ {
		topic := conf.Airports[i]

		client.Subscribe(topic, conf.Qos, onMessageReceived)
		go paho.StartSensorsPubs(client, topic, conf.Timer, conf.Qos)
	}
	<-c
}
