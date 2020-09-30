package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fri-go/internal/csv"
	"github.com/fri-go/internal/paho"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//Configuration returns a structure for the Airport name
type Configuration struct {
	Airports []string
	Timer    time.Duration
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	// process message here
	// redis
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	//csv part
	csv.GenerateCsvData(message.Payload())
	// install redis
	// configure redis (port, credentials ?, persistence)
	// tests ?
	//redis.StoreData(message.Payload())
}

func loadConfiguration() Configuration {
	//Open file
	file, _ := os.Open("config/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
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
		if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		go paho.StartSensorsPubs(client, topic, configuration.Timer)
	}

	<-c
}
