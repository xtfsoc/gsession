# gsession
A simple HTTP library


#### 需要解决的问题
- golang HTTP Response Body的内存泄漏问题
https://blog.csdn.net/hello_ufo/article/details/92994573






#### Example1:

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


#### Example:

```go
func ExampleGET() {
	session := gsession.Session()
	session.Proxy.Update("http://127.0.0.1:8888")
	session.Cookie.Add(map[string]string{"name":"wanghui", "age":"24"})

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	header["accept-encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"


	resp, err := session.GET("https://www.zhihu.com/", header, true, (1 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	//_, _ = session.GET("https://www.baidu.com/", header, true)

	fmt.Println(resp.Text())
	fmt.Println(session.Cookie.GetMap())
	//_, _ = session.GET(options)
}

func ExamplePOST() {
	session := gsession.Session()
	session.Proxy.Update("")
	session.Cookie.Add(map[string]string{"name":"wanghui", "age":"24"})

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	header["accept-encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"

	body := strings.NewReader("{\"sensor_data\":\"7a74G7m23Vrp0o5c9094051.421,2,-94,-118,199217-1,2,-94,-121,;1;2;0\"}")
	resp, err := session.POST("https://www.baidu.com/", header, body, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Content())
	fmt.Println(resp.Text()[0:10])
	fmt.Println(resp.GetCookies())

	resp, err = session.GET("https://www.baidu.com/", header, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Content())
	fmt.Println(resp.Text()[0:10])
	fmt.Println(resp.GetCookies())
}
```