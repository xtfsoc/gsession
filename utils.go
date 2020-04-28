package gsession

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

// General processing header information
// Accept-Encoding: brotli is not supported yet
// to remove it, follow the order principle br-> gzip-> deflate
func processHeader(header map[string]string) map[string]string {
	var key string
	for k, _ := range header {
		if strings.EqualFold(k, "accept-encoding") {
			key = k
		}
	}
	delete(header, key)
	header["Accept-Encoding"] = "gzip, deflate"

	return header
}

// The response cookie is automatically added to the session
func setCookie(c []*http.Cookie) {
	for _, v := range c {
		// v: *http.cookie
		// v.Value
		// v.Domain
		//cookiej[v.Name] = v.Value
		cookieSync.Store(v.Name, v.Value)
	}
}

// Processing timeout parameter, default 60s
func processTimeout(ts []time.Duration) (time.Duration, error) {
	if len(ts) == 0 {
		return 1 * time.Minute, nil
	} else if len(ts) == 1 {
		return ts[0], nil
	}
	return 0 * time.Minute, errors.New("There are multiple timeout variable parameters, should be 1")
}
