package examples_test

import (
	"fmt"
	"gsession"
	"log"
	"testing"
)

func TestFetchUO(t *testing.T) {
	session := gsession.Session()
	session.Proxy.Update("http://127.0.0.1:8888")
	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["Access-Control-Request-Method"] = "DELETE"
	header["Origin"] = "https://booking.hkexpress.com"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.128 Safari/537.36"
	header["Access-Control-Request-Headers"] = "apikey,content-type"
	header["Accept"] = "*/*"
	header["Referer"] = "https://booking.hkexpress.com/zh-CN/search/"
	header["Accept-Encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	header["Host"] = "booking-api.hkexpress.com"
	url := "https://booking-api.hkexpress.com/api/v1.0/booking/session"
	response, err := session.OPTIONS(url, header, true)
	if err != nil {
		log.Fatal(err)
	}

	header = make(map[string]string)
	header["Host"] = "booking-api.hkexpress.com"
	header["Connection"] = "keep-alive"
	header["Accept"] = "application/json, text/plain, */*"
	header["Origin"] = "https://booking.hkexpress.com"
	header["apiKey"] = "1a739d42f96658378a0ac7804fefdb2ebd649182e4971c99a3edd1e949277270"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.128 Safari/537.36"
	header["Referer"] = "https://booking.hkexpress.com/zh-CN/select/"
	header["Accept-Encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	response, err = session.DELETE(url, header, true)
	if err != nil {
		log.Fatal(err)
	}

	header = make(map[string]string)
	header["Connection"] = "keep-alive"
	header["Access-Control-Request-Method"] = "DELETE"
	header["Origin"] = "https://booking.hkexpress.com"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.128 Safari/537.36"
	header["Access-Control-Request-Headers"] = "apikey,content-type"
	header["Accept"] = "*/*"
	header["Referer"] = "https://booking.hkexpress.com/zh-CN/search/"
	header["Accept-Encoding"] = "gzip, deflate, br"
	header["Accept-Language"] = "zh-CN,zh;q=0.9,en;q=0.8"
	header["Host"] = "ancillaries-api.hkexpress.com"
	url = "https://ancillaries-api.hkexpress.com/api/v1.0/insurance/clear"
	response, err = session.OPTIONS(url, header, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Text())
}
