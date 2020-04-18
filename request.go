package gsession

import (
	"net/http"
	netUrl "net/url"
	"time"
)

//func (g gsessionObject) GET(o Options) (Response, error) {
//	o = processHeader(o)
//
//	c := &http.Client{}
//	if o.Timeout == (0 * time.Second) {
//		c.Timeout = 10 * time.Second
//	} else {
//		c.Timeout = o.Timeout
//	}
//
//	if PROXY == "" {
//
//	} else {
//		ts := &http.Transport{Proxy: func(_ *http.Request) (*url.URL, error) {
//			return url.Parse(PROXY)
//		}}
//		c.Transport = ts
//	}
//
//	req, err := http.NewRequest("GET", o.Url, nil)
//	if err != nil {
//		return nil, err
//	}
//	// 置入headers
//	for k, v := range o.Headers {
//		req.Header.Add(k, v)
//	}
//
//	// 判断是否有本地cookie
//	if len(COOKIEJ) == 0 {
//
//	} else {
//		// 有本地cookie, 自动加入
//		for k, v := range COOKIEJ {
//			var mycookie = &http.Cookie{
//				Name:  k,
//				Value: v,
//			}
//			req.AddCookie(mycookie)
//		}
//	}
//
//	response, err := c.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	cookies := response.Cookies()
//	setCookie(cookies)
//
//	var r Response
//	r = gsessionResponse{resp: response}
//	return r, nil
//}

func (g gsessionObject) GET(url string, headers map[string]string, redirect bool, timeouts ...time.Duration) (Response, error) {
	c := &http.Client{}

	// 处理各种参数
	headers = processHeader(headers)
	timeout, err := processTimeout(timeouts)
	if err != nil {
		return nil, err
	}
	c.Timeout = timeout

	if PROXY == "" {
	} else {
		ts := &http.Transport{Proxy: func(_ *http.Request) (*netUrl.URL, error) {
			return netUrl.Parse(PROXY)
		}}
		c.Transport = ts
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 置入headers
	for k, v := range headers {
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
