package gsession

import "time"

// 定义基本配置
type options struct {
	Url       string
	Mode      string
	Headers   map[string]string
	Data      string
	Redirects bool
	Timeout   time.Duration
}
