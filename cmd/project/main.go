package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fri-go/internal/csv"
	"github.com/fri-go/internal/paho"
	"github.com/fri-go/internal/redis"
	"github.com/fri-go/types/conf"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	go redis.StoreData(message.Payload())
	go csv.StoreData(message.Payload())
}

func loadConfiguration() conf.Configuration {
	file, _ := os.Open("config/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := conf.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	conf := loadConfiguration()

	client := paho.Connect(conf.Adress+":"+conf.Port, conf.ClientID)

	for i := 0; i < len(conf.Airports); i++ {
		topic := conf.Airports[i]

		client.Subscribe(topic, conf.Qos, onMessageReceived)
		go paho.StartSensorsPubs(client, topic, conf.Timer, conf.Qos)
	}
	<-c
}
