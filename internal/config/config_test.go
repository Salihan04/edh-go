package config

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestGetConfigObj(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		expected Config
	}{
		{"Valid", "../../configs/config.json", Config{
			AppClientID:    "STG2-EDH-SELF-TEST",
			PrivateKeyPath: "./ssl/STG2-EDH-SELF-TEST.pem",
			PublicCertPath: "./ssl/stg-auth-signing-public.pem",
			BaseURL:        "https://test.api.edh.gov.sg/gov/v1/entity",
			Attributes:     "entitytype,basic-profile,addresses,history,financials,capitals,declarations,charges,shareholders,appointments,licences,grants",
		}},
		{"NotExist", "notExist.json", Config{}},
		{"InvalidFormat", "invalid.json", Config{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "Valid" {
				got := GetConfigObj(tc.filename)
				if !reflect.DeepEqual(got, tc.expected) {
					t.Errorf("Got %v, expected %v", got, tc.expected)
				}
			} else if os.Getenv("WILL_EXIT") == "1" {
				GetConfigObj(tc.filename)
				return
			}
			// Test the GetConfigObj function will exit for NotExist and InvalidFormat cases
			cmd := exec.Command(os.Args[0], "-test.run=TestGetConfigObj/(NotExist|InvalidFormat)")
			cmd.Env = append(os.Environ(), "WILL_EXIT=1")
			err := cmd.Run()
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				return
			}
			t.Fatalf("process ran with err %v, want exit status 1", err)
		})
	}
}
