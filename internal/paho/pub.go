package paho

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fri-go/types/sensor"
)

func getNewData(sensorType string, idAirport string, t time.Time) sensor.Data {
	const minValue float64 = -5
	const maxValue float64 = 30
	return sensor.Data{
		IDSensor:   rand.Intn(100),
		IDAirport:  idAirport,
		SensorType: sensorType,
		Value:      minValue + rand.Float64()*(maxValue-minValue),
		Timestamp:  t,
	}
}

func dataToString(data sensor.Data) string {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(dataJSON)
}

// StartSensorsPubs : start temperature, wind and pressure sensors
func StartSensorsPubs(client mqtt.Client, topic string, period time.Duration) {
	sensors := [3]string{"temperature", "wind", "pressure"}
	timer := time.NewTicker(period * time.Second)

	for t := range timer.C {
		for _, sensor := range sensors {
			data := dataToString(getNewData(sensor, topic, t))
			client.Publish(topic, 0, false, data)
		}
	}
}
