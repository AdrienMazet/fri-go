package configuration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fri-go/types/conf"
)

// LoadConfiguration : return content of configuration file as a struct
func LoadConfiguration() conf.Configuration {
	file, _ := os.Open("config/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := conf.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
