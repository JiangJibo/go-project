package bind_func

import (
	"fmt"
	"testing"
)

// golang中给结构体或结构体指针绑定函数的区别

// 1.先创建一个结构体
type Test struct {
	name string //给结构体绑定一个字段，用以说明结构体和结构体指针绑定函数的区别
}

//2.给结构体绑定函数
func (test Test) function() {
	test.name = "结构体"
}

//3.给结构体指针绑定函数
func (test *Test) pointFunction() {
	test.name = "结构体指针"
}

func TestBindStructFunc(t *testing.T) {
	test := &Test{"创建的"}              //同Java中的new Test("创建的")，返回实例对象的引用（地址）
	test.function()                   //调用结构体函数，系统帮你转成(*test).function。
	fmt.Println("name = ", test.name) // 打印 name = "创建的"，因为是值传递，不修改name的实际值

	var test1 Test
	// java中对于只声明，没有关联内存地址的值会报空指针异常。golang不会。系统默认会转为(&test).pointFunction
	test1.pointFunction()
	fmt.Println("打印结构体指针绑定的方法")
	fmt.Println("name = ", test1.name) //name = "结构体指针",是引用传递会修改test.name的内容

}
