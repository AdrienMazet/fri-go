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

// TODO : Singleton ?
// TODO : Configure redis (port, credentials, persistence) ?
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// StoreData : store sensor data in redis database
func StoreData(message []byte) {
	var data sensor.Data
	json.Unmarshal(message, &data)
	writeData(data)
}

func writeData(data sensor.Data) {
	ISODateLayout := "2006-01-02 15:04:05"
	key := strconv.Itoa(data.IDSensor) + ":" + data.IDAirport + ":" + data.SensorType + ":" + data.Timestamp.Format(ISODateLayout)

	err := rdb.Set(ctx, key, data.Value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(key, val)
}
