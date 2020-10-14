package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

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

// GetAirportData : return all data (keys and values) for an airport
func GetAirportData(IDairport string) []sensor.Data {
	keys, err := rdb.Keys(ctx, "*:"+IDairport+":*").Result()
	if err != nil {
		fmt.Println(err)
	}

	datas := make([]sensor.Data, len(keys))

	for i := 0; i < len(keys); i++ {
		strval := rdb.Get(ctx, keys[i])
		value, err := strval.Float64()

		splittedKey := strings.Split(keys[i], ":")
		idSensor, err := strconv.Atoi(splittedKey[0])
		sensorType := splittedKey[2]
		timestamp, err := time.Parse(isoDateLayout, splittedKey[3]+":"+splittedKey[4]+":"+splittedKey[5])

		if err != nil {
			fmt.Println(err)
		}

		datas[i] = sensor.Data{
			IDSensor:   idSensor,
			IDAirport:  IDairport,
			SensorType: sensorType,
			Value:      value,
			Timestamp:  timestamp,
		}
	}

	return datas
}
