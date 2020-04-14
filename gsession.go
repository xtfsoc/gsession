package gsession

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func Request(o Options) (Response, error) {
	// 设置默认抓取超时时间
	if (0 * time.Second) == o.Timeout {
		o.Timeout = 10 * time.Second
	}
	// 判断抓取模式
	mode := strings.ToUpper(o.Mode)
	switch mode {
	case "GET":
		return get(o)
	case "POST":
		fmt.Println("POST")
	case "DELETE":
		fmt.Println("DELETE")
	default:
		return nil, errors.New("options error: 非法抓取模式(Mode)")
	}
	return nil, nil
}
