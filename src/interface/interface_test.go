package _interface

import (
	"fmt"
	"testing"
	"time"
)

// 接口方法签名
type WriteProgrammer interface {
	WriteHelloWorld() string
	//ReadHelloWorld() string  提倡小接口，否则一个struct需要实现所有接口，会很麻烦
}

type ReadProgrammer interface {
	ReadHelloWorld() string
}

// 小接口组合成大接口
type Programmer interface {
	WriteProgrammer
	ReadProgrammer
}

// 不需要显式的实现接口
type GoProgrammer struct {
}

// 没有绑定的函数, GoProgrammer 不能调用
func writeHelloWorld(s string) string {
	return "x"
}

// WriteHelloWrold() string 方法签名和 WriteProgrammer 一致
// (g *GoProgrammer) 绑定到哪个struct，结构体
func (g *GoProgrammer) WriteHelloWorld() string {
	return fmt.Sprint("Hello World")
}

func print(p GoProgrammer) string {
	p.WriteHelloWorld()
	return "ssss"
}

func TestInterface(t *testing.T) {
	g := GoProgrammer{}
	t.Log(g.WriteHelloWorld())

	// 定义接口变量
	var prog WriteProgrammer = &GoProgrammer{}
	prog.WriteHelloWorld()

	p := &GoProgrammer{}
	p.WriteHelloWorld()

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
