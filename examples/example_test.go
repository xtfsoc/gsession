package examples_test

import (
	"fmt"
	"gsession"
	"io"
	"log"
	"strings"
	"testing"
	"time"
)

func TestGsessionObject_GET(t *testing.T) {
	session := gsession.Session()
	//session.Proxy.Update("http://127.0.0.1:8888")
	session.Cookie.Update(map[string]string{"name": "wanghui", "gender": "male"})

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	header["accept-encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	resp, _ := session.GET("https://www.baidu.com/", header, true, 6*time.Second)
	fmt.Println(resp.Text())
	fmt.Println(session.Cookie.GetMap())
}

func TestGsessionObject_POST(t *testing.T) {
	session := gsession.Session()
	session.Proxy.Update("http://127.0.0.1:8888")
	session.Cookie.Update(map[string]string{"name": "wanghui", "gender": "male"})

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
	header["accept-encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"

	// Set post data
	var body io.Reader = strings.NewReader("{\"name\":\"tom\"}")
	resp, err := session.POST("https://www.baidu.com/", header, body, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Text())
}
