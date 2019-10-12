package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config is a struct that will get its values from configs/config.json
type Config struct {
	AppClientID    string `json:"app_client_id,omitempty"`
	PrivateKeyPath string `json:"private_key_path,omitempty"`
	PublicCertPath string `json:"public_cert_path,omitempty"`
	BaseURL        string `json:"base_url,omitempty"`
	Attributes     string `json:"attributes,omitempty"`
}

// GetConfig returns the JSON object in configs/config.json as a Config struct
func GetConfig(filename string) (Config, error) {
	var config Config

	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, errors.New("Unable to open config.json")
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, errors.New("Unable to unmarshal byteValue to Config struct")
	}

	return config, nil
}
