package gsession

import (
	"net/http"
)

// 获得 *http.Response 对象
type gsessionResponse struct {
	text       string
	bytes      []byte
	cookies    []*http.Cookie
	statusCode int
}

// interface接口, 模拟 gsessionResponse
type Response interface {
	Text() string
	Content() []byte
	GetCookies() []*http.Cookie
	StatusCode() int
}

func (g *gsessionResponse) Text() string {
	return g.text
	//return ConvertByte2String(b, GBK)

}

func (g *gsessionResponse) Content() []byte {
	return g.bytes
}

func (g *gsessionResponse) GetCookies() []*http.Cookie {
	return g.cookies
}

func (g *gsessionResponse) StatusCode() int {
	return g.statusCode
}
