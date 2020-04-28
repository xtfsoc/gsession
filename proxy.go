package gsession

var proxySync string

type proxy struct {
}

func (p *proxy) Update(proxy string) {
	proxySync = proxy
}
