package gsession

import (
	"net/http"
	"time"
)

// 定义基本配置
type Options struct {
	Url       string
	Mode      string
	Headers   map[string]string
	Data      string
	Proxies   string
	Redirects bool
	Timeout   time.Duration
}

// 获得 *http.Response 对象
type gsessionResponse struct {
	resp *http.Response
}

/*
cookie的生命周期管理
1. update(map[string]string) 更新cookie
2. insert(map[string]string) 增加cookie
3. delete(key) 删除单个cookie
4. clear() 清空cookie
5. getall() -> map[string]string 获取所有cookie
6. get(key) -> string 获取单个cookie值
*/

// Cookie
//type Cookie struct {
//}
//type Teacher struct {
//	Cookie //组合People
//	//People2 //组合People2
//}
//
//func (c *Cookie) Clear() {
//	fmt.Println("People2:showD")
//}

//type Cookie interface {
//	Clear() error
//}

//func (R gsessionResponse) Clear() error {
//	fmt.Println("clear Cookie")
//	return nil
//}

// interface接口, 模拟 gsessionResponse
type Response interface {
	Text() string
	Content() []uint8
}

// func (R gsessionResponse) Text() string {
