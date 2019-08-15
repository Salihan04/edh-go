package config

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
	}{
		{"NotExist", "notExist.json"},
		{"Valid", "../../configs/config.json"},
		{"MissingFields", "./test_files/missing_fields.json"},
		{"InvalidFormat", "./test_files/invalid.json"},
	}

	expected := getConfigFields(Config{
		AppClientID:    "app_client_id",
		PrivateKeyPath: "private_key_path",
		PublicCertPath: "public_cert_path",
		BaseURL:        "base_url",
		Attributes:     "attributes",
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "Valid" {
				c, _ := GetConfig(tc.filename)
				got := getConfigFields(c)
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("Got %v, expected %v", got, expected)
				}
			} else if tc.name == "MissingFields" {
				c, _ := GetConfig(tc.filename)
				got := getConfigFields(c)
				if reflect.DeepEqual(got, expected) {
					t.Errorf("There are missing config fields. Got %v, expected %v", got, expected)
				}
			} else {
				_, err := GetConfig(tc.filename)
				if err == nil {
					t.Errorf("Expected error to be returned\n")
				}
			}
		})
	}
}

func getConfigFields(c Config) []string {
	var fields []string
	e := reflect.ValueOf(&c).Elem()

	for i := 0; i < e.NumField(); i++ {
		// Only append fields with non-empty values
		if e.Field(i).Interface() != "" {
			fields = append(fields, e.Type().Field(i).Name)
		}
	}

	return fields
}
