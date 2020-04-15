package gsession

import "time"

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
