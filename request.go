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

//func (g gsessionObject) GET(o options) (Response, error) {
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
//	if len(cookiej) == 0 {
//
//	} else {
//		// 有本地cookie, 自动加入
//		for k, v := range cookiej {
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
	var c *http.Client
	if redirect {
		c = &http.Client{}
	} else {
		c = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				//return errors.New("Disable redirects")
				return http.ErrUseLastResponse
			},
		}
	}

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
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookiej.Range(f)

	if len(keys) == 0 {

	} else {
		// 有本地cookie, 自动加入
		for i := 0; i < len(keys); i++ {
			v, _ := cookiej.Load(keys[i])
			if v == nil {
				return nil, errors.New(fmt.Sprintf("添加Cookie失败, 值为空: %v\n", v))
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

	var r Response
	r = &gsessionResponse{text: string(b), bytes: b, cookies: cookies}
	return r, nil
}

func (g gsessionObject) POST(url string, headers map[string]string, body io.Reader, redirect bool, timeouts ...time.Duration) (Response, error) {

	// 处理重定向
	var c *http.Client
	if redirect {
		c = &http.Client{}
	} else {
		c = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				//return errors.New("Disable redirects")
				return http.ErrUseLastResponse
			},
		}
	}

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

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// 置入headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// 判断是否有本地cookie
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookiej.Range(f)

	if len(keys) == 0 {

	} else {
		// 有本地cookie, 自动加入
		for i := 0; i < len(keys); i++ {
			v, _ := cookiej.Load(keys[i])
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

	var r Response
	r = &gsessionResponse{text: string(b), bytes: b, cookies: cookies}
	return r, nil

}

func (g gsessionObject) PUT(o options) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) DELETE(o options) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) OPTIONS(o options) (Response, error) {
	panic("implement me")
}
