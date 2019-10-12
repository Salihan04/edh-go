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
		{"Valid", "./test_files/valid.json"},
		{"MissingFields", "./test_files/missing_fields.json"},
		{"InvalidFormat", "./test_files/invalid.json"},
	}

	expected := Config{
		AppClientID:    "app_client_id",
		PrivateKeyPath: "private_key_path",
		PublicCertPath: "public_cert_path",
		BaseURL:        "base_url",
		Attributes:     "attributes",
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "Valid" {
				got, _ := GetConfig(tc.filename)
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("got %v, expected %v", got, expected)
				}
			} else if tc.name == "MissingFields" {
				got, _ := GetConfig(tc.filename)
				t.Logf("%v, %v\n", got, expected)
				if reflect.DeepEqual(got, expected) {
					t.Errorf("there are missing config fields. Got %v, expected %v", got, expected)
				}
			} else {
				_, err := GetConfig(tc.filename)
				if err == nil {
					t.Errorf("expected error to be returned\n")
				}
			}
		})
	}
}
