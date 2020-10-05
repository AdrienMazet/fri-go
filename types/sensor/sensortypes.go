package sensor

import (
	"time"
)

// Data : data send by sensor
type Data struct {
	IDSensor   int
	IDAirport  string
	SensorType string
	Value      float64
	Timestamp  time.Time
}
