package gsession

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

/*
通用处理头部信息
1. Accept-Encoding: 暂不支持br, 要把它拿掉, 按照顺位原则br -> gzip -> deflate
*/
func processHeader(header map[string]string) map[string]string {

	var key string
	for k, _ := range header {
		if strings.EqualFold(k, "accept-encoding") {
			key = k
		}
	}
	delete(header, key)
	header["Accept-Encoding"] = "gzip, deflate"

	return header
}

/*
通用setCookie
抓取后的setCookie保存到COOKIEJ
*/
func setCookie(c []*http.Cookie) {
	for _, v := range c {
		// v: *http.cookie
		// v.Value
		// v.Domain
		COOKIEJ[v.Name] = v.Value
	}
}

/*
处理超时参数
*/
func processTimeout(ts []time.Duration) (time.Duration, error) {
	if len(ts) == 0 {
		return 1 * time.Minute, nil
	} else if len(ts) == 1 {
		return ts[0], nil
	}
	return 0 * time.Minute, errors.New("There are multiple timeout variable parameters, should be 1")
}
