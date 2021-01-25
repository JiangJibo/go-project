package groutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func run(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	//ch := make(chan string)
	// buffer channel，这样不会阻塞携程往里面放数据, 协程就不会阻塞而一直存在
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	t.Log("After:", runtime.NumGoroutine())
}
