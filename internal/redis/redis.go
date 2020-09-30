package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fri-go/types/sensor"
	"github.com/go-redis/redis/v8"
)

// faire un singleton ?

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// StoreData : store sensor data in redis database
func StoreData(data sensor.Data) {
	key := strconv.Itoa(data.IDSensor) + ":" + data.IDAirport + ":" + data.SensorType + data.Timestamp.String()
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
