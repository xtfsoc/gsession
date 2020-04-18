package gsession

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
func (c *cookie) GetMap() map[string]string {
	return COOKIEJ
}

/*
添加cookie, 等同于更新cookie
*/
func (c *cookie) Add(cookie map[string]string) {
	COOKIEJ = cookie
}

/*
更新cookie
原则: 已有的key覆盖掉, 换成新的值
没有的值加上
*/
func (c *cookie) Update(cookie map[string]string) {
	COOKIEJ = cookie
}

/*
清空cookie
*/
func (c *cookie) Clear() {
	COOKIEJ = make(map[string]string)
}
