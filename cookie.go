package gsession

import (
	"fmt"
	"sync"
)

/*
cookie的生命周期管理
1. update(map[string]string) 更新cookie
2. insert(map[string]string) 增加cookie
3. delete(key) 删除单个cookie
4. clear() 清空cookie
5. getall() -> map[string]string 获取所有cookie
6. get(key) -> string 获取单个cookie值
*/

var cookiej sync.Map

type cookie struct{}

/*
获取全部cookies
*/
func (c *cookie) GetMap() sync.Map {
	return cookiej
}

/*
添加cookie, 等同于更新cookie
sync.Map是覆盖增
*/
func (c *cookie) Add(cookie map[string]string) {
	for k, v := range cookie {
		cookiej.Store(k, v)
	}
	//cookiej = cookie
}

/*
更新cookie
原则: 已有的key覆盖掉, 换成新的值
没有的值加上
*/
func (c *cookie) Update(cookie map[string]string) {
	for k, v := range cookie {
		fmt.Println("添加cookie", k, v)
		cookiej.Store(k, v)
	}
}

/*
清空cookie
*/
func (c *cookie) Clear() {
	//cookiej = make(map[string]string)
	var keys []string
	f := func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		return true
	}
	cookiej.Range(f)

	for i := 0; i < len(keys); i++ {
		cookiej.Delete(keys[i])
	}
}
