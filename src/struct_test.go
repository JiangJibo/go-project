package src

import (
	"fmt"
	"testing"
	"unsafe"
)

type Person struct {
	Id   string
	Name string
	Age  int
}

//func (e *Person) String() string {
//	fmt.Printf("Address is %d\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func (e Person) String() string {
	fmt.Printf("Address is %d\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	p := Person{"1", "bob", 15}
	fmt.Printf("Address is %d\n", unsafe.Pointer(&p.Name))
	t.Log(p.String())

}
