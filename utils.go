package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Prints to stdout
func p(a ...interface{}) {
	fmt.Println(a)
}

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}

var config Config

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Cannot decode configuration from file", err)
	}
}

