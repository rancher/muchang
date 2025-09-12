package dara

import (
	"net/url"
	"strings"
)

// PercentEncode encodes a string for use in URLs, replacing certain characters.
func PercentEncode(uri string) string {
	if uri == "" {
		return ""
	}
	uri = url.QueryEscape(uri)
	uri = strings.Replace(uri, "+", "%20", -1)
	uri = strings.Replace(uri, "*", "%2A", -1)
	uri = strings.Replace(uri, "%7E", "~", -1)
	return uri
}
