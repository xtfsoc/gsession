package gsession

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
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

// The response cookie is automatically added to the SessionObject
func setCookie(c []*http.Cookie) {
	for _, v := range c {
		// v: *http.cookie
		// v.Value
		// v.Domain
		// cookiej[v.Name] = v.Value
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

// General request method
func requestHTTP(url string, mode string, headers map[string]string, body io.Reader, redirect bool, timeouts []time.Duration) (Response, error) {
	var c *http.Client

	// Set redirect
	if redirect {
		c = &http.Client{}
	} else {
		c = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				// handle redirect
				// return errors.New("Disable redirects")
				return http.ErrUseLastResponse
			},
		}
	}

	// Set timeout
	timeout, err := processTimeout(timeouts)
	if err != nil {
		return nil, err
	}
	c.Timeout = timeout

	// Set Proxy
	if proxySync != "" {
		ts := &http.Transport{Proxy: func(_ *http.Request) (*netUrl.URL, error) {
			return netUrl.Parse(proxySync)
		}}
		c.Transport = ts
	}

	req, err := http.NewRequest(mode, url, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	headers = processHeader(headers)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Set cookies
	// Determine if there is a local cookie
	var keys []string
	// Traverse all the keys of cookieSync and put them in the keys
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookieSync.Range(f)
	// local cookies, automatically add
	if len(keys) > 0 {
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			var v interface{}
			var ok bool = false
			for !ok {
				v, ok = cookieSync.Load(k)
			}
			if v == nil {
				return nil, errors.New(fmt.Sprintf("Failed to add cookie, the value is empty: %v\n", v))
			}
			req.AddCookie(&http.Cookie{Name: keys[i], Value: v.(string)})
		}
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	// Add response cookie to cookieSync
	cookies := resp.Cookies()
	setCookie(cookies)

	var reader io.ReadCloser
	var encode = resp.Header.Get("Content-Encoding")
	if strings.Contains(strings.ToLower(encode), "gzip") {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = resp.Body
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var r Response
	r = &gsessionResponse{text: string(b), bytes: b, cookies: cookies, statusCode: resp.StatusCode}
	return r, nil
}
