package requester

import (
	"edh-go/internal/config"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var appID, clientID, attributes string
var txnNo, timestamp int64

func init() {
	c, err := config.GetConfig("../../configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	appID = c.AppClientID
	clientID = c.AppClientID
	attributes = c.Attributes
	txnNo = time.Now().Unix()
	timestamp = time.Now().Unix() * 1000
}

func formulateBaseString(httpMethod string, url string, appID string,
	attributes string, clientID string, nonce string,
	timestamp int64, txnNo int64) string {
	baseParams := map[string]string{
		"app_id":           appID,
		"attributes":       attributes,
		"client_id":        clientID,
		"nonce":            nonce,
		"signature_method": "RS256",
		"timestamp":        strconv.FormatInt(timestamp, 10),
		"txn_no":           strconv.FormatInt(txnNo, 10),
	}
	var keys []string

	for k := range baseParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	baseString := fmt.Sprintf("%v&%v", strings.ToUpper(httpMethod), url)
	for _, k := range keys {
		baseString += fmt.Sprintf("&%v=%v", k, baseParams[k])
	}

	return baseString
}

func formulateURLWithQueryString(url string) string {
	return fmt.Sprintf("%v?attributes=%v&client_id=%v&txnNo=%v", url, attributes, clientID, txnNo)
}
