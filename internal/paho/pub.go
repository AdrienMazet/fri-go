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

func getNewData() {
	// faire une factory a data
	//return new data()
}

// config from file

func startPubTemp(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := data{
			idSensor:   01,
			idAirport:  "CDG",
			sensorType: "temperature",
			value:      (-5) + rand.Float64()*(30-(-5)), //valeur a modifier
			timestamp:  t,
		}
		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
		//client.Publish(topic, 0, false, d1)
	}
}

func startPubWind(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := data{
			idSensor:   01,
			idAirport:  "JFK",
			sensorType: "vent",
			value:      (0) + rand.Float64()*(100-0), //valeur a modifier
			timestamp:  t,
		}
		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
		//client.Publish(topic, 0, false, "wind : "+t.String())
	}
}

func startPubPressure(client mqtt.Client, topic string, period time.Duration) {
	timer := time.NewTicker(period * time.Second)
	for t := range timer.C {
		d1 := data{
			idSensor:   01,
			idAirport:  "JFK",
			sensorType: "pression",
			value:      (1) + rand.Float64()*(5-1), //valeur a modifier
			timestamp:  t,
		}
		//conversion en string sinon peut pas envoyer de message
		text := fmt.Sprintf(
			"idSensor:" + strconv.Itoa(d1.idSensor) + "\n" +
				"idAirport:" + d1.idAirport + "\n" +
				"sensorType:" + d1.sensorType + "\n" +
				"value:" + strconv.FormatFloat(d1.value, 'f', -1, 64) + "\n" +
				"timestamp:" + d1.timestamp.Format("2006-01-02")) //on doit donner un exemple de formattage pour qu'il comprend
		client.Publish(topic, 0, false, text)
		//client.Publish(topic, 0, false, "pressure : "+t.String())
	}
}

// StartSensorsPubs : start all three sensors
func StartSensorsPubs(client mqtt.Client, topic string, period time.Duration) {
	rand.Seed(time.Now().UnixNano())
	go startPubTemp(client, topic, period)
	go startPubWind(client, topic, period)
	go startPubPressure(client, topic, period)
}
