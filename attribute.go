package gsession

import (
	"net/http"
	"time"
)

// 定义基本配置
type Options struct {
	Url       string
	Mode      string
	Headers   map[string]string
	Data      string
	Proxies   string
	Redirects bool
	Timeout   time.Duration
}

// 获得 *http.Response 对象
type response struct {
	resp *http.Response
}

// interface接口, 模拟 response
type Response interface {
	Text() string
	Content() []uint8
}
