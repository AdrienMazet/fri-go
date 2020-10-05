package paho

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type data struct {
	idSensor   int
	idAirport  string
	sensorType string
	value      float64
	timestamp  time.Time
}

func getNewData(sensorType string, idAirport string, t time.Time) data {
	return data{
		idSensor:   01,
		idAirport:  idAirport,
		sensorType: sensorType,
		value:      (-5) + rand.Float64()*(30-(-5)), //valeur a modifier
		timestamp:  t,
	}
}

func dataToText(d1 data) string {
	return fmt.Sprintf(
		"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
			"idAirport:" + d1.idAirport + "\n" +
			"sensorType:" + d1.sensorType + "\n" +
			"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
			"timestamp:" + d1.timestamp.Format("2006-01-02")) //we must give on example to the Format function to make it work
}

// StartSensorsPubs : start all three sensors
func StartSensorsPubs(client mqtt.Client, topic string, period time.Duration) {
	sensors := [3]string{"temperature", "wind", "pressure"}
	timer := time.NewTicker(period * time.Second)

	for t := range timer.C {
		for _, sensor := range sensors {
			data := dataToText(getNewData(sensor, topic, t))
			client.Publish(topic, 0, false, data)
		}
	}
}
