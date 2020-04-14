package gsession

import (
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8     = Charset("UTF-8")
	GB18030  = Charset("GB18030")
	GBK      = Charset("GBK")
	HZGB2312 = Charset("HZGB2312")
)

// 调用 fmt.Printf("%s\n", ConvertByte2String(stdoutStderr, GB18030))
func ConvertByte2String(b []byte, charset Charset) string {
	var s string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	case UTF8:
		fallthrough
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	case HZGB2312:
		var decodeBytes, _ = simplifiedchinese.HZGB2312.NewDecoder().Bytes(b)
		s = string(decodeBytes)
	default:
		s = string(b)
	}

	return s
}

/*
版权声明：本文为CSDN博主「追随她的旅程」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/erecrlion/article/details/88632707
result = ConvertToString(html, "gbk", "utf-8")
*/
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
