package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Config struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

var once = &sync.Once{}
var config = &Config{}

func main() {
	for {
		conf := GetConfig()
		fmt.Printf("%+v\n", conf)
		time.Sleep(time.Second)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		fmt.Println("Load config from file")
		marshalled, err := ioutil.ReadFile("./config.json")
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(marshalled, &config); err != nil {
			panic(err)
		}

	})

	return config
}
