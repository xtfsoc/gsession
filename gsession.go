package gsession

import "time"

type Options struct {
	Url       string
	Mode      string
	Headers   map[string]string
	Data      string
	Proxies   string
	Redirects bool
	Timeout   time.Duration
}

func init() {
	COOKIEJ = make(map[string]string)
}

// gsessionResponse
type gessionObject struct {
	//mysession session
}

// Response
type GsessionObject interface {
	GET(o Options) (Response, error)
	POST(o Options) (Response, error)
	PUT(o Options) (Response, error)
	DELETE(o Options) (Response, error)
	OPTIONS(o Options) (Response, error)
	//GetAllCookies() map[string]string

}

func Session() GsessionObject {
	var obj GsessionObject
	obj = gessionObject{}
	return obj
}
