package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fri-go/types/sensor"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// return filename like : idAirport_timestamp_sensorType.csv prefixed with the data lake folder
func generateFilename(data sensor.Data) string {
	folder := "data_lake/"
	ISODateLayout := "2006-01-02"
	return folder + data.IDAirport + "_" + data.Timestamp.Format(ISODateLayout) + "_" + data.SensorType + ".csv"
}

//create the csv file or add to the end of the existing file
func appendDataToCsv(data sensor.Data) {
	filename := generateFilename(data)
	fileExist := fileExists(filename)

	if !fileExist {
		// create file
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println("Can't create file :"+filename, err)
		}
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExist {
		// write headers
		writer.Write([]string{"IDSensor", "Value"})
	}

	writer.Write([]string{strconv.Itoa(data.IDSensor), strconv.FormatFloat(data.Value, 'E', -1, 64)})

	if err := writer.Error(); err != nil {
		fmt.Println(err)
	}
}

// StoreData : add data from sensors to csv
func StoreData(message []byte) {
	var data sensor.Data
	json.Unmarshal(message, &data)
	appendDataToCsv(data)
}
