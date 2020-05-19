package examples_test

import (
	"compress/gzip"
	"fmt"
	"gsession"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGoutineGession(t *testing.T) {
	for i := 0; i < 100; i++ {
		go requestGession(i)
	}
	time.Sleep(1 * time.Hour)
}

func requestGession(sign int) {
	for {
		session := gsession.Session()
		// session.Proxy.Update("http://127.0.0.1:8888")
		session.Cookie.Update(map[string]string{"name": "wanghui", "age": "22"})

		header := make(map[string]string)
		header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.3"
		header["accept-encoding"] = "gzip, deflate, br"
		header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
		header["Host"] = "www.zhihu.com"

		resp, err := session.GET("https://www.zhihu.com/", header, true, (100 * time.Second))
		if err != nil {
			fmt.Printf("*ERROR*: %v\n", err)
			continue
		}
		_, err = fmt.Printf("sign: %d, statusCode: %d, text: %s\n", sign, resp.StatusCode(), resp.Text()[100:170])
	}

}

func requestHTTP(sign int) {
	url := "https://www.zhihu.com/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Connection", "close")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cache-control", "no-cache")

	for {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("*ERROR*: %v\n", err)
			continue
		}
		if res != nil {
			defer res.Body.Close()
		}

		var reader io.ReadCloser
		var encode = res.Header.Get("Content-Encoding")
		if strings.Contains(strings.ToLower(encode), "gzip") {
			reader, _ = gzip.NewReader(res.Body)
		} else {
			reader = res.Body
		}

		defer func() {
			fmt.Println("start======================")
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("end=========================")
		}()

		b, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Printf("*ERROR*: %v\n", err)
			continue
		}

		fmt.Printf("sign: %d, statusCode: %d, text: %s\n", sign, res.StatusCode, string(b)[100:170])
	}

}

func TestIOutil(t *testing.T) {
	var reader io.ReadCloser
	fmt.Println(reader)
	_, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
}
