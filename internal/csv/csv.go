package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func generateFilename(tableau []string) string {
	//create the filename : idAirport_timestamp_sensorType.csv
	//TODO : travailler avec l'objet Data et pas un tableau de string

	fmt.Println(tableau)
	// folder where the CSV are stored
	folder := "data_lake/"
	//idAirport_timestamp_sensorType.csv
	return folder + (strings.Split(tableau[1], ":")[1] + "_" +
		strings.Split(tableau[4], "timestamp:")[1] + "_" +
		strings.Split(tableau[2], ":")[1] + ".csv")
}

//format the data before adding it to the csv file
func formatData(tableau []string, header bool) [][]string {
	//TODO : travailler avec l'objet Data et pas un tableau de string

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
func appendDataToCsv(tab []string) {
	filename := generateFilename(tab)
	data := formatData(tab, false)

	if !fileExists(filename) {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatal("Can't create file :"+filename, err)
		}
		data = formatData(tab, true)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(data)

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending succed")
}

// StoreData add data from sensors to csv
func StoreData(message []byte) {
	// TODO : byte[] to string to obj Data (in method ?)

	//string conversion used to filter data
	res := string(message)
	//split by newline and parsing using the function
	appendDataToCsv(strings.Split(res, "\n"))
}
