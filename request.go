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

	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	cookies := response.Cookies()
	setCookie(cookies)

	var r Response
	r = gsessionResponse{resp: response}
	return r, nil
}

/*
通用setCookie
*/
func setCookie(c []*http.Cookie) {

	for _, v := range c {
		// v: *http.Cookie
		//fmt.Println("v.Value:", v.Value)
		//fmt.Println("v.Domain:", v.Domain) // string
		//fmt.Println("v.MaxAge:", v.MaxAge) // int
		COOKIEJ[v.Name] = v.Value
	}
}
