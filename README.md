# gsession
A simple HTTP library


#### 需要解决的问题

- golang HTTP Response Body的内存泄漏问题
https://blog.csdn.net/hello_ufo/article/details/92994573






#### Example:

```go
//options := gsession.Options{}
//options.Url = "https://www.baidu.com"
//options.Mode = "get"
//options.Timeout = 40 * time.Second
header := make(map[string]string)
header["Connection"] = "keep-alive"
header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"
header["Accept-Encoding"] = "gzip, deflate, br"
header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
//options.Headers = header
//options.Proxies = ""
//options.Proxies = "http://127.0.0.1:8888"

options := gsession.Options{
    Url:       "https://www.baidu.com",
    Mode:      "get",
    Headers:   header,
    Data:      "",
    Proxies:   "",
    Redirects: true,
    Timeout:   40 * time.Second,
}

//resp, _ := gsession.Request(options)
var s = gsession.Session()
//s.Cookie.GetAll()
resp, _ := s.GET(options)

fmt.Println(resp.Text())

fmt.Println(s.Cookie.GetAll())

//var s = gsession.Session111()
//_, _ = s.GET(options)
```


加入代理新操作方式

```go
session := gsession.Session()
session.Proxy.Update("htpsss:////sdsdsdd")
session.Cookie.GetAll()
```


加版本