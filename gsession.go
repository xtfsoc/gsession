package gsession

import (
	"io"
	"time"
)

type SessionObject struct {
	Proxy  proxy
	Cookie cookie
	gsessionFunc
}

type gsessionObject struct{}

type gsessionFunc interface {
	GET(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	POST(url string, headers map[string]string, body io.Reader, redirect bool, timeout ...time.Duration) (Response, error)
	PUT(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	DELETE(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	HEAD(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	OPTIONS(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	CONNECT(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	TRACE(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
}

func Session() SessionObject {
	// 新增session要清空COOKIEJ
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookieSync.Range(f)

	for i := 0; i < len(keys); i++ {
		cookieSync.Delete(keys[i])
	}

	sessionInit := func() gsessionFunc {
		var ga gsessionFunc
		ga = gsessionObject{}
		return ga
	}
	// return SessionObject{proxy{}, cookie{}, sessionInit()}
	return SessionObject{
		Proxy:        proxy{},
		Cookie:       cookie{},
		gsessionFunc: sessionInit(),
	}
}
