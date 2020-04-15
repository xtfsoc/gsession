package gsession

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"strings"
)

// 获取报文的文本内容
func (R gsessionResponse) Text() string {
	r := R.resp
	defer r.Body.Close()

	var reader io.ReadCloser
	if strings.Contains(strings.ToLower(r.Header.Get("Content-Encoding")), "gzip") {
		reader, _ = gzip.NewReader(r.Body)
	} else {
		reader = r.Body
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}
	//s := string(b)
	//return ConvertToString(s, "utf-8", "utf-8")
	//return ConvertByte2String(b, HZGB2312)
	//return mahonia.NewEncoder("utf-8").ConvertString(s)
	return string(b)
}

// 获得字节流报文
func (R gsessionResponse) Content() []uint8 {
	r := R.resp
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []uint8("")
	}
	return b
}
