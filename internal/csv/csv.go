package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

//test if the file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//generate the filename
func generateFilename(tableau []string) string {
	// folder where the CSV are stored
	folder := "data_lake/"
	//idAirport_timestamp_sensorType.csv
	return folder + (strings.Split(tableau[1], ":")[1] + "_" +
		strings.Split(tableau[4], "timestamp:")[1] + "_" +
		strings.Split(tableau[2], ":")[1] + ".csv")
}

//format the data before adding it to the csv file
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

// GenerateCsvData generate csv from sensors data
func GenerateCsvData(message []byte) {
	//string conversion used to filter data
	res := string(message)
	//split by newline and parsing using the function
	generateDataLake(strings.Split(res, "\n"))
}
