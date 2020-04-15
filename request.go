package gsession

import (
	"net/http"
	"net/url"
	"time"
)

func (g gsob) GET(o Options) (Response, error) {
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

	// 判断是否有本地cookie
	if len(COOKIEJ) == 0 {

	} else {
		// 有本地cookie, 自动加入
		for k, v := range COOKIEJ {
			var mycookie = &http.Cookie{
				Name:  k,
				Value: v,
			}
			req.AddCookie(mycookie)
		}
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

func (g gsob) POST(o Options) (Response, error) {
	panic("implement me")
}

func (g gsob) PUT(o Options) (Response, error) {
	panic("implement me")
}

func (g gsob) DELETE(o Options) (Response, error) {
	panic("implement me")
}

func (g gsob) OPTIONS(o Options) (Response, error) {
	panic("implement me")
}

/*
通用setCookie
*/
func setCookie(c []*http.Cookie) {
	for _, v := range c {
		// v: *http.Cookie
		// v.Value/v.Domain
		COOKIEJ[v.Name] = v.Value
	}
}
