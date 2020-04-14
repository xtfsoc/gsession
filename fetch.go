package gsession

import (
	"net/http"
	"net/url"
	"time"
)

func get(o Options) (Response, error) {
	c := &http.Client{}
	if o.Timeout == (0 * time.Second) {
		c.Timeout = 10 * time.Second
	} else {
		c.Timeout = o.Timeout
	}

	if o.Proxies == "" {
	} else {
		ts := &http.Transport{Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse(o.Proxies)
		}}
		c.Transport = ts
	}

	req, err := http.NewRequest("GET", o.Url, nil)
	if err != nil {
		return nil, err
	}
	// 置入headers
	for k, v := range o.Headers {
		req.Header.Add(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var r Response
	r = response{resp: resp}
	return r, nil
}
