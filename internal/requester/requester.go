package requester

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func formulateBaseString(httpMethod string, url string, appID string, attributes string, clientID string, nonce string, timestamp int64, txnNo int64) string {
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
