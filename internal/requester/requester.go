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

type fixedParams struct {
	url       string
	nonce     string
	txnNo     int64
	timestamp int64
}

var c config.Config
var p fixedParams

func init() {
	var err error
	c, err = config.GetConfig("../../configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	p = fixedParams{
		url:       c.BaseURL, //TODO: set to new value when UEN is available
		nonce:     "",        //TODO: call a function to generate nonce
		txnNo:     time.Now().Unix(),
		timestamp: time.Now().Unix() * 1000,
	}
}

func formulateBaseString(httpMethod string, p fixedParams, c config.Config) string {
	baseParams := map[string]string{
		"app_id":           c.AppClientID,
		"attributes":       c.Attributes,
		"client_id":        c.AppClientID,
		"nonce":            p.nonce,
		"signature_method": "RS256",
		"timestamp":        strconv.FormatInt(p.timestamp, 10),
		"txn_no":           strconv.FormatInt(p.txnNo, 10),
	}
	var keys []string

	for k := range baseParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	baseString := fmt.Sprintf("%v&%v", strings.ToUpper(httpMethod), p.url)
	for _, k := range keys {
		baseString += fmt.Sprintf("&%v=%v", k, baseParams[k])
	}

	return baseString
}

func formulateURLWithQueryString(p fixedParams, c config.Config) string {
	return fmt.Sprintf("%v?attributes=%v&client_id=%v&txnNo=%v", p.url, c.Attributes, c.AppClientID, p.txnNo)
}
