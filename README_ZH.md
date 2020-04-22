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



#### Examples

* [Example GET](#Example GET)



----

#### Example GET

> 完整的GET请求方式, 以baidu.com为例.



- 简单请求

  ```go
  session := gsession.Session()
  
  header := make(map[string]string)
  header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
  header["accept-encoding"] = "gzip, deflate, br"
  header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
  
  resp, err := session.GET("https://www.zhihu.com/", header, true, (1 * time.Second))
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println(resp.Text())
  ```

- 完整请求

  ```go
  // 获取session对象
  session := gsession.Session()
  
  // 不启用代理的方式:
  // 1. session.Proxy.Update("")
  // 2. 直接注释掉
  session.Proxy.Update("http://127.0.0.1:8888")
  
  // 自定义添加Cookie
  // 添加Cookie方式:
  // 1. session.Cookie.Add(map[string]string{"name":"tom", "age":"24"})
  // 2. header["Cookie"] = "name=tom; age=24"
  session.Cookie.Add(map[string]string{"name":"tom", "age":"24"})
  
  header := make(map[string]string)
  header["Connection"] = "keep-alive"
  header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
  header["accept-encoding"] = "gzip, deflate, br"
  header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
  
  // GET方法参数
  // session.GET(url string, headers map[string]string, redirect bool, timeout ...time.Duration)
  // url: 请求网址
  // header: 请求头部信息
  // redirect: 是否重定向
  // timeout: 超时时间, 不写默认60s
  resp, err := session.GET("https://www.zhihu.com/", header, true, (1 * time.Second))
  if err != nil {
  	log.Fatal(err)
  }
  
  // 获取当前session的全部代理
  fmt.Println(session.Cookie.GetMap())
  
  // 再次请求, session会自动把上一步链接的setCookie合并请求
  // 如果不想合并请求, 使用 session.Cookie.Clear()
  session.Cookie.Add(map[string]string{"name":"tom", "age":"24"})
  _, _ = session.GET("https://www.baidu.com/", header, true)
  
  fmt.Println(resp.Text())
  fmt.Println(resp.Content())
  fmt.Println(resp.GetCookies())
  ```

