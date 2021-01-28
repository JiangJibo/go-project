package channel

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cannelChan chan int) bool {
	select {
	case <-cannelChan: // close 会使所有处理等待状态的case返回, 得到的是channel的零值; 从关闭了的channel里获取数据时可以的，只是得到的是零值
		fmt.Println("true")
		return true
	default:
		fmt.Println("false")
		return false
	}
}

func cancel_1(cancelChan chan int) {
	cancelChan <- 1
}

func cancel_2(cancelChan chan int) {
	// close 会使所有处理等待状态的case返回, 得到的是channel的零值
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan int)
	for i := 0; i < 5; i++ {
		go func(x int, cancelChan chan int) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(x, "Done")
		}(i, cancelChan)
	}
	time.Sleep(time.Nanosecond * 1)
	cancel_2(cancelChan)
	b := <-cancelChan
	fmt.Println("get value from cancelled channel:  ", b)
	time.Sleep(time.Millisecond * 100)
}
