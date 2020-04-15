package gsession

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// 获得 *http.Response 对象
type gsessionResponse struct {
	resp *http.Response
}

// interface接口, 模拟 gsessionResponse
type Response interface {
	Text() string
	Content() []uint8
}

func (g gsessionResponse) Text() string {
	r := g.resp
	defer r.Body.Close()

	var reader io.ReadCloser
	if strings.Contains(strings.ToLower(r.Header.Get("Content-Encoding")), "gzip") {
		reader, _ = gzip.NewReader(r.Body)
	} else {
		reader = r.Body
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}
	return string(b)
}

func (g gsessionResponse) Content() []uint8 {
	r := g.resp
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []uint8("")
	}
	return b
}
