package gsession

import "io/ioutil"

// 获取报文的文本内容
func (R response) Text() string {
	r := R.resp
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ""
	}
	s := string(b)
	return s
}

// 获得字节流报文
func (R response) Content() []uint8 {
	r := R.resp
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []uint8("")
	}
	return b
}
