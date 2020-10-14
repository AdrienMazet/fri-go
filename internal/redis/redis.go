package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fri-go/types/sensor"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var isoDateLayout string = "2006-01-02 15:04:05"

// StoreData : store sensor data in redis database
func StoreData(message []byte) {
	var data sensor.Data
	json.Unmarshal(message, &data)
	writeData(data)
}

func writeData(data sensor.Data) {
	key := strconv.Itoa(data.IDSensor) + ":" + data.IDAirport + ":" + data.SensorType + ":" + data.Timestamp.Format(isoDateLayout)

	err := rdb.Set(ctx, key, data.Value, 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetAverageSensorValue : get average value for a sensor of an airport on a date
func GetAverageSensorValue(idAirport string, date string, sensorType string) float64 {
	keys, err := rdb.Keys(ctx, "*:"+idAirport+":"+sensorType+":"+date+"*").Result()
	if err != nil {
		fmt.Println(err)
	}

	values := make([]float64, len(keys))

	for i := 0; i < len(keys); i++ {
		strval := rdb.Get(ctx, keys[i])
		value, err := strval.Float64()
		if err != nil {
			fmt.Println(err)
		}

		values[i] = value
	}

	var total float64 = 0
	for _, value := range values {
		total += value
	}
	average := total / float64(len(values))

	return average
}
