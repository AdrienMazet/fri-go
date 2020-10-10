package models

import (
	"fmt"
    "time"
	"github.com/go-redis/redis/v8"
)

const layoutISO = "2006-01-02"

type WeatherData struct {
	key   string `json:"key"`
	value string `json:"value"`
}

type DataList []WeatherData

client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

ctx := client.Context()

func filterDataByDate(dateA string, dateB string) {
    startDate, _ := time.Parse(layoutISO, dateA)
    endDate, _ := time.Parse(layoutISO, dateB)
    currentDate := startDate
    for (currentDate.Before(endDate)) {
        val, err := client.Get(ctx, dateA).Result()
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(val);
        currentDate = currentDate.Add(24 * time.Hour)
    }

}

func getAverageValues() {

}
