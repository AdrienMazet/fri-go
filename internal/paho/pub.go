package paho

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fri-go/types/sensor"
)

func getNewData(sensorType string, idAirport string, t time.Time) sensor.Data {
	return sensor.Data{
		IDSensor:   01,
		IDAirport:  idAirport,
		SensorType: sensorType,
		Value:      (-5) + rand.Float64()*(30-(-5)), //valeur a modifier
		Timestamp:  t,
	}
}

func dataToText(data sensor.Data) string {
	return fmt.Sprintf(
		"IDSensor:" + strconv.Itoa(data.IDSensor) + "\n" +
			"IDAirport:" + data.IDAirport + "\n" +
			"SensorType:" + data.SensorType + "\n" +
			"Value:" + strconv.FormatFloat(data.Value, 'f', -1, 64) + "\n" +
			"Timestamp:" + data.Timestamp.Format("2006-01-02")) //we must give on example to the Format function to make it work
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
