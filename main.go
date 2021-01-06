package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config is ...
type Config struct {
	Token string `json:"token"`
	Tall  int    `json:"tall"`
}

var jsonFile string = "./.config"
var config Config = loadConfiguration(jsonFile)

func main() {
	fmt.Println("Token string:", config.Token)
	fmt.Println("Tall int:", config.Tall)
}

func loadConfiguration(file string) Config {
	var config Config

	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(configFile, &config)

	log.Println(file, "is loaded and ready.")
	return config
}
