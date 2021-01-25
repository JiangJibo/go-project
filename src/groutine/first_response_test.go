package groutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
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
	// 当channel里有一条数据时立即返回，10个携程有一个写进去就返回了

	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	t.Log("After:", runtime.NumGoroutine())
}
