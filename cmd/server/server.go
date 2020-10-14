package main

import (
	"fmt"

	"github.com/fri-go/internal/configuration"
	"github.com/fri-go/internal/redis"
)

func main() {
	conf := configuration.LoadConfiguration()

	for i := 0; i < len(conf.Airports); i++ {
		airport := conf.Airports[i]
		fmt.Println(redis.GetAirportData(airport))
		fmt.Println(".")
	}
}
