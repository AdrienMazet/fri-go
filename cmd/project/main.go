package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fri-go/internal/csv"
	"github.com/fri-go/internal/paho"
	"github.com/fri-go/types/conf"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func dataLakeHandler(client mqtt.Client, message mqtt.Message) {
	csv.GenerateCsvData(message.Payload())
}

func redisHandler(client mqtt.Client, message mqtt.Message) {
	// install redis
	// configure redis (port, credentials ?, persistence)
	// tests ?
	//redis.StoreData(message.Payload())
}

func loadConfiguration() conf.Configuration {
	//Open file
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

	client := paho.Connect("tcp://localhost:1883", "mqtt_client")

	//Load configuration
	configuration := loadConfiguration()
	//Sub to each topic from each airport
	for i := 0; i < len(configuration.Airports); i++ {
		topic := configuration.Airports[i]

		// redis sub
		if token := client.Subscribe(topic, 0, redisHandler); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		// data lake sub
		if token := client.Subscribe(topic, 0, dataLakeHandler); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		go paho.StartSensorsPubs(client, topic, configuration.Timer)
	}

	<-c
}
