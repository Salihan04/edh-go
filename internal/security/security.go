package security

import (
	"errors"
	"io/ioutil"
)

func loadPrivateKey(path string) (string, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.New("unable to open private key file at " + path)
	}
	return string(dat), nil
}

// SignBaseString signs a base string with a given private key and returns the signature
func SignBaseString(privKey string, baseString string) string {
	return ""
}
