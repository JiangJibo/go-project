package error

import (
	"errors"
	"fmt"
	"testing"
)

// 异常恢复 os.Exit 不会打印堆栈, 不会执行defer的方法
func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong"))
}
