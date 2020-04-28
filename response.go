package gsession

import "net/http"

type gsessionResponse struct {
	text       string
	bytes      []byte
	cookies    []*http.Cookie
	statusCode int
}

type Response interface {
	Text() string
	Content() []byte
	GetCookies() []*http.Cookie
	StatusCode() int
}

func (g *gsessionResponse) Text() string {
	// return ConvertByte2String(b, GBK)
	return g.text
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
