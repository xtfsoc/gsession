package gsession

import (
	"io"
	"time"
)

type session struct {
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
	OPTIONS(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
}

// Exported function with unexported return type
// Inspection info: Reports exported functions with unexported return types.
// Unexported types can be difficult to use when viewing documentation under go doc
func Session() session {
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
	// return session{proxy{}, cookie{}, sessionInit()}
	return session{
		Proxy:        proxy{},
		Cookie:       cookie{},
		gsessionFunc: sessionInit(),
	}
}
