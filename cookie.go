package gsession

import "fmt"

/*
cookie的生命周期管理
1. update(map[string]string) 更新cookie
2. insert(map[string]string) 增加cookie
3. delete(key) 删除单个cookie
4. clear() 清空cookie
5. getall() -> map[string]string 获取所有cookie
6. get(key) -> string 获取单个cookie值
*/

//var COOKIEJ map[string]string

var COOKIEJ map[string]string

func GetAllCookie() map[string]string {
	return COOKIEJ
}

// 我的尝试
type Cookie struct{}

func (c *Cookie) GetAll() {
	fmt.Println("GetALL")
}

type Session1 struct {
	Cookie //组合People
}

func Session() Session1 {
	return Session1{}
}

//type cookiej struct {
//	cj map[string]string
//}

//type resp1 struct {
//	Cookie
//}
//
//func (c *Cookie) Clear() {
//	fmt.Println("showB")
//}
