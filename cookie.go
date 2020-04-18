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

var COOKIEJ map[string]string

type cookie struct{}

/*
获取全部cookies
*/
func (c *cookie) GetAll() map[string]string {

	fmt.Println("GetALL:", PROXY)
	return COOKIEJ
}

/*
更新cookie
原则: 已有的key覆盖掉, 换成新的值
没有的值加上
*/
func (c *cookie) Update(ckj map[string]string) {
	COOKIEJ = ckj
}
