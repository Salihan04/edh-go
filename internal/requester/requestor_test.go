package requester

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestFormulateBaseString(t *testing.T) {
	testCases := []struct {
		name       string
		httpMethod string
		url        string
		appID      string
		attributes string
		clientID   string
		nonce      string
		timestamp  int64
		txnNo      int64
		expected   string
	}{
		{
			"Valid",
			"Get",
			"https://test.api.edh.gov.sg/gov/v1/entity/201800001A",
			"STG2-EDH-SELF-TEST",
			"basic-profile,addresses",
			"STG2-EDH-SELF-TEST",
			"1234567890abcde",
			123,
			123,
			"GET&https://test.api.edh.gov.sg/gov/v1/entity/201800001A&" +
				"app_id=STG2-EDH-SELF-TEST&attributes=basic-profile,addresses&client_id=STG2-EDH-SELF-TEST&" +
				"nonce=1234567890abcde&signature_method=RS256&timestamp=123&txn_no=123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := formulateBaseString(tc.httpMethod, tc.url, tc.appID,
				tc.attributes, tc.clientID, tc.nonce,
				tc.timestamp, tc.txnNo)
			baseStringParts := strings.Split(got, "&")

			httpMethod := baseStringParts[0]
			if !reflect.DeepEqual(httpMethod, "GET") {
				t.Errorf("got %v, expected %v", httpMethod, "GET")
			}

			a := strings.Split(baseStringParts[3], "=")[1]
			attributes := strings.Split(a, ",")
			expectedAttributes := []string{"basic-profile", "addresses"}
			if !reflect.DeepEqual(attributes, expectedAttributes) {
				t.Errorf("attributes should be comma separated string")
			}

			nonce := strings.Split(baseStringParts[5], "=")[1]
			if !reflect.DeepEqual(len(nonce), 15) {
				t.Errorf("nonce must be of length 15")
			}

			signatureMethod := strings.Split(baseStringParts[6], "=")[1]
			if !reflect.DeepEqual(signatureMethod, "RS256") {
				t.Errorf("signature method must be RS256")
			}

			timestamp := strings.Split(baseStringParts[7], "=")[1]
			if _, err := strconv.ParseInt(timestamp, 10, 64); err != nil {
				t.Errorf("timestamp must be an integer")
			}

			txnNo := strings.Split(baseStringParts[8], "=")[1]
			if _, err := strconv.ParseInt(txnNo, 10, 64); err != nil {
				t.Errorf("timestamp must be an integer")
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}

func TestFormulateURLWithQueryString(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected string
	}{
		{
			"Valid",
			"https://test.api.edh.gov.sg/gov/v1/entity/201800001A",
			"https://test.api.edh.gov.sg/gov/v1/entity/201800001A?" +
				"attributes=" + attributes + "&" +
				"client_id=STG2-EDH-SELF-TEST&" +
				"txnNo=" + strconv.FormatInt(txnNo, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := formulateURLWithQueryString(tc.url)
			if got != tc.expected {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
