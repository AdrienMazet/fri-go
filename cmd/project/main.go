package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"../../internal/paho"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//Configuration returns a structure for the Airport name
type Configuration struct {
	Airports []string
	Timer    time.Duration
}

//test if the file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//function to generate the filename
func generateFilename(tableau []string) string {
	// folder where the CSV are stored
	folder := "data_lake/"
	//idAirport_timestamp_sensorType.csv
	return folder + (strings.Split(tableau[1], ":")[1] + "_" +
		strings.Split(tableau[4], "timestamp:")[1] + "_" +
		strings.Split(tableau[2], ":")[1] + ".csv")
}

//function to format the data before adding it to the csv file
func formatData(tableau []string, header bool) [][]string {
	var data [][]string
	if header {
		data = append(data, []string{"idSensor", "Valeur"})
	}
	//we split on the ':' to retrieve the second value (actual one)
	//tableau[0] = idSensor
	//tableau[3] = value
	data = append(data, []string{strings.Split(tableau[0], ":")[1], strings.Split(tableau[3], ":")[1]})
	return (data)
}

//create the csv file or add to the end of the existing file
func generateDataLake(tab []string) {
	//create the filename : idAirport_timestamp_sensorType.csv
	filename := generateFilename(tab)
	//test is the file exists and append data or create it and append data
	if fileExists(filename) {
		//open the file and append data
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		//format data before appending
		data := formatData(tab, false)
		//write data to file
		w := csv.NewWriter(f)
		w.WriteAll(data)

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Appending succed")
	} else {
		//create the file and append data
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal("Impossible de créer le fichier", err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		//format the data to append
		data := formatData(tab, true)
		for _, value := range data {
			err := writer.Write(value)
			if err != nil {
				log.Fatal("Impossible d'écrire dans le fichier", err)
			}
		}
		fmt.Println("Creation & Appending succed")
	}
}

func generateCsvData(message []byte) {
	//string conversion used to filter data
	res := string(message)
	//split by newline and parsing using the function
	generateDataLake(strings.Split(res, "\n"))
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	// process message here
	// redis
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	//csv part
	generateCsvData(message.Payload())
}

func loadConfiguration() Configuration {
	//Open file
	file, _ := os.Open("config/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	client := paho.Connect("tcp://localhost:1883", "mqtt_client")

	//Load configuration
	configuration := loadConfiguration()
	//Sub to each topic from each airport
	for i := 0; i < len(configuration.Airports); i++ {
		topic := configuration.Airports[i]
		if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		paho.StartSensorsPubs(client, topic, configuration.Timer)
	}

	<-c
}
