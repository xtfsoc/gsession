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

type Cookie struct{}

/*
获取全部cookies
*/
func (c *Cookie) GetAll() map[string]string {
	return COOKIEJ
}

/*
增加cookie
*/
