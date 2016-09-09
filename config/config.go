package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Configuration exported struct to allow existinece in main
type Configuration struct {
	feeds []string
}

//ParseConfig Exported function to parse config file
func ParseConfig(path string, Config *Configuration) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could Not open Config File")
	}
	_, err = toml.DecodeReader(file, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
