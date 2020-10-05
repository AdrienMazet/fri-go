package conf

import "time"

//Configuration : configuration for airports publications
type Configuration struct {
	Airports []string
	Timer    time.Duration
}
