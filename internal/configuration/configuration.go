package configuration

import (
	"encoding/json"
<<<<<<< HEAD
<<<<<<< HEAD
	"log"
=======
	"fmt"
>>>>>>> d455e4a... moved server folder in internal
=======
	"log"
>>>>>>> 9b86a94... refacto
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
<<<<<<< HEAD
		log.Fatal(err)
=======
		fmt.Println("error:", err)
>>>>>>> d455e4a... moved server folder in internal
=======
		log.Fatal(err)
>>>>>>> 9b86a94... refacto
	}
	return configuration
}
