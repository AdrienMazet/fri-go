package conf

import "time"

//Configuration : configuration for airports publications
type Configuration struct {
	Adress   string
	Port     string
	ClientID string
	Qos      byte
	Airports []string
	Timer    time.Duration
}
