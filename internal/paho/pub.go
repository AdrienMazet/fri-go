package paho

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// func getNewData() sensor.Data {
// 	// faire une factory a data
// 	//return new data()
// 	return nil
// }

func startPubTemp(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, "temp : "+t.String())
	}
}

func startPubWind(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, "wind : "+t.String())
	}
}

func startPubPressure(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, "pressure : "+t.String())
	}
}

// StartSensorsPubs : start all three sensors
func StartSensorsPubs(client mqtt.Client, topic string, period time.Duration) {
	go startPubTemp(client, topic, period)
	go startPubWind(client, topic, period)
	go startPubPressure(client, topic, period)
}
