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

// config from file
func startPubTemp(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := getNewData("temperature", topic, t)

		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
	}
}

func startPubWind(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := getNewData("vent", topic, t)

		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
	}
}

func startPubPressure(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := getNewData("pression", topic, t)

		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
	}
}

// StartSensorsPubs : start all three sensors
func StartSensorsPubs(client mqtt.Client, topic string, period time.Duration) {
	rand.Seed(time.Now().UnixNano())
	go startPubTemp(client, topic, period)
	go startPubWind(client, topic, period)
	go startPubPressure(client, topic, period)
}
