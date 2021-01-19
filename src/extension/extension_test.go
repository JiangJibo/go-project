package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Printf("pet speak to %s", host)
}

type Dog struct {
	Pet // 匿名嵌套对象，获得了Pet的方法,类似继承; 但不支持方法重写，也就是Dog不能重写Pet方法
}

func (d *Dog) Speak() {
	fmt.Print("Wang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("duck")
}

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
