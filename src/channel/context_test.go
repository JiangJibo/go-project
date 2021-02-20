package channel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelledWithContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		//fmt.Println("true")
		return true
	default:
		//fmt.Println("false")
		return false
	}
}

// 通过Context来终止任务
func TestCancelWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(x int, ctx context.Context) {
			for {
				if isCancelledWithContext(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(x, "Done")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Millisecond * 100)
}

func TestName(t *testing.T) {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}
