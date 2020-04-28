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

func (g gsessionObject) POST(url string, headers map[string]string, body io.Reader, redirect bool, timeout ...time.Duration) (Response, error) {
	var c *http.Client
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

	// Process parameters
	headers = processHeader(headers)
	tm, err := processTimeout(timeout)
	if err != nil {
		return nil, err
	}
	c.Timeout = tm

	if proxySync == "" {
	} else {
		ts := &http.Transport{Proxy: func(_ *http.Request) (*netUrl.URL, error) {
			return netUrl.Parse(proxySync)
		}}
		c.Transport = ts
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// Determine if there is a local cookie
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookieSync.Range(f)

	if len(keys) == 0 {

	} else {
		// local cookies, automatically add
		for i := 0; i < len(keys); i++ {
			v, _ := cookieSync.Load(keys[i])
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

	cookies := resp.Cookies()
	setCookie(cookies)

	var reader io.ReadCloser
	var encode = resp.Header.Get("Content-Encoding")
	if strings.Contains(strings.ToLower(encode), "gzip") {
		reader, _ = gzip.NewReader(resp.Body)
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