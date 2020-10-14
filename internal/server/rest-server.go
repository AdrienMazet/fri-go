package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	router := InitializeRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
