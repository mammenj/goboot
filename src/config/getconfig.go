package config

import (
	"encoding/json"
	"os"
)

func GetConfiguration() (Configuration, error) {
	myconfig := Configuration{}
	file, err := os.Open("C:\\Users\\john.mammen01\\GoglandProjects\\goboot2\\goboot\\src\\config\\config.json")
	if err != nil {
		return myconfig, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&myconfig)
	if err != nil {
		return myconfig, err
	}
	return myconfig, nil
}
