package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config is a struct that will get its values from configs/config.json
type Config struct {
	AppClientID    string `json:"app_client_id"`
	PrivateKeyPath string `json:"private_key_path"`
	PublicCertPath string `json:"public_cert_path"`
	BaseURL        string `json:"base_url"`
	Attributes     string `json:"attributes"`
}

// GetConfigObj returns the JSON object in configs/config.json as a Config struct
func GetConfigObj(filename string) Config {
	var config Config

	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open config.json: %v\n", err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Unable to unmarshal byteValue to Config struct: %v\n", err)
	}

	return config
}
