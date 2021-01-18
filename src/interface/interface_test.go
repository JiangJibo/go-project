package _interface

import (
	"fmt"
	"testing"
	"time"
)

// 接口方法签名
type Programmer interface {
	WriteHelloWrold() string
}

// 不需要显式的实现接口
type GoProgrammer struct {
}

// 没有绑定的函数, GoProgrammer 不能调用
func writeHelloWorld(s string) string {
	return "x"
}

// WriteHelloWrold() string 方法签名和 Programmer 一致
// (g *GoProgrammer) 绑定到哪个struct，结构体
func (g *GoProgrammer) WriteHelloWrold() string {
	return fmt.Sprint("Hello World")
}

func print(p GoProgrammer) string {
	p.WriteHelloWrold()
	return "ssss"
}

func TestInterface(t *testing.T) {
	g := GoProgrammer{}
	t.Log(g.WriteHelloWrold())

	// 定义接口变量
	var prog Programmer = &GoProgrammer{}
	prog.WriteHelloWrold()

	p := &GoProgrammer{}
	p.WriteHelloWrold()

}

// 自定义类型
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}
