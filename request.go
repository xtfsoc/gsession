package gsession

import (
	"io"
	"time"
)

func (g gsessionObject) GET(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "GET", headers, nil, redirect, timeout)
}

func (g gsessionObject) POST(url string, headers map[string]string, body io.Reader, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "POST", headers, body, redirect, timeout)
}

func (g gsessionObject) PUT(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "PUT", headers, nil, redirect, timeout)
}

func (g gsessionObject) DELETE(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "DELETE", headers, nil, redirect, timeout)
}

func (g gsessionObject) HEAD(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "HEAD", headers, nil, redirect, timeout)
}

func (g gsessionObject) OPTIONS(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "OPTIONS", headers, nil, redirect, timeout)
}

func (g gsessionObject) CONNECT(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "CONNECT", headers, nil, redirect, timeout)
}

func (g gsessionObject) TRACE(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	return requestBase(url, "TRACE", headers, nil, redirect, timeout)
}
