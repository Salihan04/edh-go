package config

import (
	"os"
	"os/exec"
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
				got := getConfigFields(GetConfig(tc.filename))
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("Got %v, expected %v", got, expected)
				}
			} else if tc.name == "MissingFields" {
				got := getConfigFields(GetConfig(tc.filename))
				if !reflect.DeepEqual(got, expected) {
					t.Logf("Missing config fields. Got %v, expected %v", got, expected)
				}
			} else if os.Getenv("WILL_EXIT") == "1" {
				GetConfig(tc.filename)
				return
			}
			// Test the GetConfig function will exit for NotExist and InvalidFormat cases
			simulateExit(t)
		})
	}
}

func simulateExit(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestGetConfig/(NotExist|InvalidFormat)")
	cmd.Env = append(os.Environ(), "WILL_EXIT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
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
