package gsession

import (
	"io"
	"time"
)

// 初始化
func init() {
	COOKIEJ = make(map[string]string)
}

type session struct {
	Proxy  proxy
	Cookie cookie
	gsessionAction
}

func Session() session {
	// 新增session要清空COOKIEJ
	COOKIEJ = make(map[string]string)

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
