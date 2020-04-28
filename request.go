package gsession

import (
	"time"
)

func (g gsessionObject) PUT(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) DELETE(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	panic("implement me")
}

func (g gsessionObject) OPTIONS(url string, headers map[string]string, redirect bool, timeout ...time.Duration) (Response, error) {
	panic("implement me")
}
