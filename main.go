package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config is ...
// Struct må være tilpasset for innholdet i JSON filen. Viktig å bruke riktig tick (`) type.
// Første kommentar over struct definisjon må se ut som den gjør ovenfor.
// Navnene i struct på starte med stor bokstav.
type Config struct {
	Token string `json:"token"`
	Tall  int    `json:"tall"`
}

// Setter .config til jsonFile variable, gjør det globalt i tilfelle
// de skal brukes i flere funksjoner.
var jsonFile string = "./.config"

// Oppretter en Config struct basert på jsonFile
var config Config = loadConfiguration(jsonFile)

// Enkel main som bare printer til console verdiene fra .config
func main() {
	fmt.Println("Token string:", config.Token)
	fmt.Println("Tall int:", config.Tall)
}

// Funksjon for å laste inn .config filen i en Config struct
func loadConfiguration(file string) Config {
	// Definerer en ny Config instans
	var config Config

	// Leser innhold fra .config filen til configFile som []byte
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}

	// "Pakker ut" data fra configFile til config struct
	json.Unmarshal(configFile, &config)

	fmt.Println(file, "is loaded and ready.")

	// returnerer config struct, klar til bruk.
	return config
}
