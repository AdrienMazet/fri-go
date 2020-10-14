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

// Average : average values of 3 sensors
type Average struct {
	Temperature float64
	Wind        float64
	Pressure    float64
}
