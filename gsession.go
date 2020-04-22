package gsession

import (
	"io"
	"time"
)

// 初始化
//func init() {
//	cookiej = make(map[string]string)
//}

type session struct {
	Proxy  proxy
	Cookie cookie
	gsessionAction
}

func Session() session {
	// 新增session要清空COOKIEJ
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookiej.Range(f)

	for i := 0; i < len(keys); i++ {
		cookiej.Delete(keys[i])
	}

	sessionInit := func() gsessionAction {
		var ga gsessionAction
		ga = gsessionObject{}
		return ga
	}
	return session{proxy{}, cookie{}, sessionInit()}
}

type gsessionAction interface {
	GET(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error)
	POST(url string, headers map[string]string, body io.Reader, redirect bool, timeout ...time.Duration) (Response, error)
	PUT(o options) (Response, error)
	DELETE(o options) (Response, error)
	OPTIONS(o options) (Response, error)
	// GetAllCookies() map[string]string
}

type gsessionObject struct{}

//func sessionInit() gsessionAction {
//	var ga gsessionAction
//	ga = gsessionObject{}
//	return ga
//}
