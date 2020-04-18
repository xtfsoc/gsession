package gsession_test

import (
	"awesomeProject1/gsession"
	"fmt"
)

func typeA() {
	session := gsession.Session()
	session.Proxy.Update("http://127.0.0.1:8888")
	cookie := make(map[string]string)
	cookie["name"] = "wanghui"
	cookie["age"] = "24"
	session.Cookie.Update(cookie)

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	header["accept-encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"

	resp, _ := session.GET("https://www.baidu.com/", header, true)
	_, _ = session.GET("https://www.baidu.com/", header, true)

	fmt.Println(resp.Text())
	fmt.Println(session.Cookie.GetMap())
	//_, _ = session.GET(options)
}

/*
此使用方式lastest
commitid: fa409c8ebcf625128ef6e67d88df1f146cbba27a
message: 加版本
*/
func typeB() {
	//header := make(map[string]string)
	//header["Connection"] = "keep-alive"
	//header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	//header["accept-encoding"] = "gzip, deflate, br"
	//header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	//options := gsession.Options{
	//	Url:       "https://www.zhihu.com",
	//	Mode:      "get",
	//	Headers:   header,
	//	Data:      "",
	//	Redirects: true,
	//	Timeout:   40 * time.Second,
	//}

	//var s = gsession.Session()
	//resp, _ := s.GET(options)
	//
	//fmt.Println(resp.Text())
	//fmt.Println("------")
	//fmt.Println(s.Cookie.GetMap())
}
