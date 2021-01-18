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
	Pet // 匿名嵌套对象，获得了Pet的方法,类似继承; 但不支持方法重载，也就是Dog不能重载Pet方法
}

func (d *Dog) Speak() {
	fmt.Print("Wang")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("duck")
}
