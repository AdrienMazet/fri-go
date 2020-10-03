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

type Configuration struct {
	Airports []string
	Timer    time.Duration
}

//teste si un fichier existe
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//cree le fichier csv ou ajoute les donnees a la fin
func generateDataLake(tab []string) {
	//creer le nom du fichier : idAirport_timestamp_sensorType.csv
	filename := strings.Split(tab[1], ":")[1] + "_" +
		strings.Split(tab[4], "timestamp:")[1] + "_" +
		strings.Split(tab[2], ":")[1] + ".csv"
	//tester si le csv existe avec le nom du fichier
	if fileExists(filename) {
		//ouvrir le csv et ajouter les donnees
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		//formattage des donnees avant ajout
		var data [][]string
		data = append(data, []string{strings.Split(tab[0], ":")[1], strings.Split(tab[3], ":")[1]})
		//ecrire les donnees
		w := csv.NewWriter(f)
		w.WriteAll(data)

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Appending succed")
	} else {
		//creer le fichier et ajouter les donnees
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal("Impossible de créer le fichier", err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		//formatter les donnees a ajouter
		var data [][]string
		data = append(data, []string{"idSensor", "Valeur"})
		data = append(data, []string{strings.Split(tab[0], ":")[1], strings.Split(tab[3], ":")[1]})
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
	//conversion en string pour pouvoir filtrer
	res := string(message)
	//split sur les retours charriots puis parsing avec la fonction
	generateDataLake(strings.Split(res, "\n"))
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	// process message here
	// redis
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	//partie csv
	generateCsvData(message.Payload())
}

func loadConfiguration() Configuration {
	//Open file
	file, _ := os.Open("conf.json")
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
