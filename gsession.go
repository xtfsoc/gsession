package gsession

import (
	"fmt"
	"io"
	"time"
)

func init() {
	COOKIEJ = make(map[string]string)
	fmt.Println("gsession默认init()方法")
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
	PUT(o Options) (Response, error)
	DELETE(o Options) (Response, error)
	OPTIONS(o Options) (Response, error)
	// GetAllCookies() map[string]string
}

type gsessionObject struct{}

//func sessionInit() gsessionAction {
//	var ga gsessionAction
//	ga = gsessionObject{}
//	return ga
//}
