package main

import (
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"net/http"

	"github.com/fri-go/internal/server"
)

func main() {
	fmt.Println("Server listening on port 8080")
	fmt.Println("2 routes availables :")
	fmt.Println("GET /{idAirport}/{sensorType}/filter/{date1}/{date2}")
	fmt.Println("GET /{idAirport}/{date}/results/average")

	router := server.InitializeRouter()
	fmt.Println(http.ListenAndServe(":8080", router))
=======
=======
	"log"
=======
>>>>>>> 9b86a94... refacto
	"net/http"
>>>>>>> 756fc63... launch server and create routes

	"github.com/fri-go/internal/server"
)

func main() {
	fmt.Println("Server listening on port 8080")
	fmt.Println("2 routes availables :")
	fmt.Println("GET /{idAirport}/{sensorType}/filter/{date1}/{date2}")
	fmt.Println("GET /{idAirport}/{date}/results/average")

<<<<<<< HEAD
	for i := 0; i < len(conf.Airports); i++ {
		airport := conf.Airports[i]
		fmt.Println(redis.GetAirportData(airport))
		fmt.Println(".")
	}
>>>>>>> d455e4a... moved server folder in internal
=======
	router := server.InitializeRouter()
<<<<<<< HEAD
	log.Fatal(http.ListenAndServe(":8080", router))
>>>>>>> 756fc63... launch server and create routes
=======
	fmt.Println(http.ListenAndServe(":8080", router))
>>>>>>> 9b86a94... refacto
}
