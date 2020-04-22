# gsession
简单, 稳定的HTTP请求库.

实现OPTIONS/GET/HEAD/POST/PUT/DELETE/TRACE/CONNECT八种请求方式. 



#### Installation

```
go get -u github.com/wanghuijz/gsession
```



#### Usage

```go
import "github.com/wanghuijz/gsession"
```





----

#### Examples

###### GET

> 获取session, 请求网址, 以baidu.com为例

```go
session := gsession.Session()
// 启用代理的方式:
session.Proxy.Update("http://127.0.0.1:8888")

// 自定义添加Cookie
session.Cookie.Add(map[string]string{"name":"tom", "age":"24"})

header := make(map[string]string)
header["Connection"] = "keep-alive"
header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
header["accept-encoding"] = "gzip, deflate, br"
header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"

// session.GET 参数, timeout可以不写, 默认60s
// GET(url string, headers map[string]string, redirect bool, timeout ...time.Duration)
resp, err := session.GET("https://www.zhihu.com/", header, true, (1 * time.Second))
if err != nil {
	log.Fatal(err)
}
fmt.Println(session.Cookie.GetMap())

session.Cookie.Add(map[string]string{"name":"tom", "age":"24"})
_, _ = session.GET("https://www.baidu.com/", header, true)

fmt.Println(resp.Text())
fmt.Println(resp.Content())
fmt.Println(resp.GetCookies())
```

