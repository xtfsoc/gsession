package gsession

/*
Proxy 模块
*/

var PROXY string

type proxy struct {
}

func (p *proxy) Update(proxy string) {
	PROXY = proxy
}
