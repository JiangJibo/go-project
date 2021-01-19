package _interface

import (
	"fmt"
	"testing"
)

// 空接口
func DoSomething(p interface{}) {
	/*if _, ok := p.(int); ok {
		fmt.Println("Integer")
	} else if _, ok := p.(string); ok {
		fmt.Println("string")
	} else {
		fmt.Println("Unknown type")
	}*/
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknown type", v)
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething('x')
}
