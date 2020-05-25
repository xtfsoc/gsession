package gsession

var proxySync string

type proxy struct{}

// example: Update("http://127.0.0.1:8888")
func (p *proxy) Update(proxy string) {
	proxySync = proxy
}
