package gsession

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (g gsessionObject) GET(o Options) (Response, error) {
	o = processHeader(o)

	c := &http.Client{}
	if o.Timeout == (0 * time.Second) {
		c.Timeout = 10 * time.Second
	} else {
		c.Timeout = o.Timeout
	}

	if PROXY == "" {

	} else {
		ts := &http.Transport{Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse(PROXY)
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

func (g gsessionObject) POST(o Options) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) PUT(o Options) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) DELETE(o Options) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) OPTIONS(o Options) (Response, error) {
	panic("implement me")
}

/*
通用setCookie
*/
func setCookie(c []*http.Cookie) {
	for _, v := range c {
		// v: *http.cookie
		// v.Value/v.Domain
		COOKIEJ[v.Name] = v.Value
	}
}

/*
通用处理头部信息
1. Accept-Encoding: 暂不支持br, 要把它拿掉, 按照顺位原则br -> gzip -> deflate
*/
func processHeader(o Options) Options {
	var header map[string]string = o.Headers
	var key string
	for k, _ := range header {
		if strings.EqualFold(k, "accept-encoding") {
			key = k
		}
	}
	delete(header, key)
	header["Accept-Encoding"] = "gzip, deflate"

	o.Headers = header
	return o
}
