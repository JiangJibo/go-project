package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 包装函数, 类似装饰者模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFt(t *testing.T) {
	var a, _ = returnMultiValues()
	t.Log(a)
	timeSpentFunc := timeSpent(slowFunc)
	x := timeSpentFunc(1)
	t.Log(x)
}

// 可变参数
func SumVarPrams(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParams(t *testing.T) {
	sum := SumVarPrams(1, 2, 3, 4)
	t.Log(sum)
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}()
	t.Log("Started")
	panic("Fatal error")
	t.Log("Closed")
}
