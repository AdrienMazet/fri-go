package configuration

import (
	"encoding/json"
<<<<<<< HEAD
	"log"
=======
	"fmt"
>>>>>>> d455e4a... moved server folder in internal
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
<<<<<<< HEAD
		log.Fatal(err)
=======
		fmt.Println("error:", err)
>>>>>>> d455e4a... moved server folder in internal
	}
	return configuration
}
