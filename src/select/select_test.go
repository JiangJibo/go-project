package _select

import (
	"fmt"
	"testing"
	"time"
)

func service(n int) string {
	time.Sleep(time.Millisecond * time.Duration(n))
	fmt.Printf("---service---立即执行 %d 毫秒延时\n", n)
	return "---service---Done"
}

// chan >> channel
func AsyncService(n int) chan string { // 该函数返回一个通道, 消息类型是 string
	//retCh := make(chan string) // 声明一个无缓冲 chan通道, 这个通道的类型是 string, 通道用于协程间的通信
	// // ⚠️ 使用无缓冲通道时, 如果协程1 没有立即接受协程 2 通过管道发送的消息，就会阻塞，反之亦然
	retCh := make(chan string, 2) // 声明一个缓冲容量为 1 的chan通道, 这个通道的类型是 string, 通道用于协程间的通信
	go func() {                   // 此处执行 匿名协程函数
		ret := service(n)
		fmt.Println("---AsyncService---returned result:", ret)
		retCh <- ret //通过通道将 service()函数返回的值传递给 通道 retCh
		fmt.Println("---AsyncService---service exited")
	}()
	return retCh // 返回这个通道
}

func TestSelect(t *testing.T) {
	select {
	case ret1 := <-AsyncService(1000):
		t.Logf(ret1)
	case ret2 := <-AsyncService(1300):
		t.Logf(ret2)
	case <-time.After(time.Millisecond * 2000):
		t.Error("time out")
	}
}
