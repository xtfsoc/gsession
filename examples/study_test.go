package examples_test

import (
	"fmt"
	"testing"
)

// person结构体，包含年龄，名称，车
type Person struct {
	age  int
	name string
	car  Car
}

// person名下的车
type Car struct {
	// 车的名字
	name string
}

// 一个存放person的map
var personMap map[string]Person

// 给参数person设置名字
func setName(person Person, name string) {
	person.name = name
}

// 设置名字
func (person Person) setName(name string) {
	person.name = name
}

// 打印person的名字
func printName(person Person) {
	fmt.Println(person.name)
}

// 结构体person自己支持打印名字
func (person Person) printName() {
	fmt.Println(person.name)
}

func TestZhizhen(t *testing.T) {
	person := Person{}
	fmt.Println(person) //{0  {}}
	person.age = 12
	person.name = "小明"
	person.car = Car{"宝马"}
	fmt.Println(person) //{12 小明 {宝马}}，正常赋值给person变量，因为这是在方法里面的变量
	setName(person, "小红")
	fmt.Println(person) //{12 小明 {宝马}}，小红赋值失败，传递给setName方法的person没有赋值成功
	person.setName("小红")
	fmt.Println(person) //{12 小明 {宝马}}，person自己setName，还是失败
	personMap = make(map[string]Person)
	personMap["test"] = person
	person = personMap["test"]
	person.name = "小红"
	fmt.Println(person)               //{12 小红 {宝马}},从map中取出person，给小红赋值成功
	for _, value := range personMap { //遍历map
		fmt.Println(value) //{12 小明 {宝马}}，打印的还是小明，而不是小红，说明上面personMap["test"]对象赋值失败
	}
}
