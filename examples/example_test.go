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

func TestNetHttp(t *testing.T) {
	fmt.Println("OPOPOP")
	n, err := division(1, 0)
	if err != nil {
		fmt.Printf("捕获到[001]: %v\n", err)
	}
	fmt.Println(n)
	fmt.Println("assssssssas")
}

func division(x, y int) (result int, err error) {
	defer func() {
		if e := recover(); e != nil {
			// 这句话的作用是: e是interface, 需要转成error类型
			// 然后把参数err赋值e, 最后统一return出去
			err = e.(error)
		}
	}()
	result = x / y
	return result, nil
}

func TestChaifen(t *testing.T) {
	fmt.Println("aaaa")
	defer func() {
		if e := recover(); e != nil {
			// 这句话的作用是: e是interface, 需要转成error类型
			// 然后把参数err赋值e, 最后统一return出去
			fmt.Println("bbbb")
			fmt.Println(e)
		}
	}()
	fmt.Println("cccc")
	x := 9
	y := 0
	result := x / y
	fmt.Println("dddd")
	fmt.Println(result)
}

func TestS(t *testing.T) {
	var s = " Get "
	s = strings.Trim(s, " ")
	s = strings.ToUpper(s)
	fmt.Println(s)
}
