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
	return sensor.Data{
		IDSensor:   01, // TODO : randomize this
		IDAirport:  idAirport,
		SensorType: sensorType,
		Value:      (-5) + rand.Float64()*(30-(-5)), // ??
		Timestamp:  t,                               // TODO : format
	}
}

func dataToText(data sensor.Data) string {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	return string(dataJSON)
	// return fmt.Sprintf(
	// 	"IDSensor:" + strconv.Itoa(data.IDSensor) + "\n" +
	// 		"IDAirport:" + data.IDAirport + "\n" +
	// 		"SensorType:" + data.SensorType + "\n" +
	// 		"Value:" + strconv.FormatFloat(data.Value, 'f', -1, 64) + "\n" +
	// 		"Timestamp:" + data.Timestamp.Format("2006-01-02")) //we must give on example to the Format function to make it work
}

// StartSensorsPubs : start temperature, wind and pressure sensors
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
