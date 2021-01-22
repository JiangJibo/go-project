package groutine

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("---service---立即执行 50 毫秒延时")
	time.Sleep(time.Millisecond * 50)
	return "---service---Done"
}

func otherTask() {
	fmt.Println("---otherTask---working on something else")
	fmt.Println("---otherTask---立即执行 100 毫秒延时")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("---otherTask---Task is Done")
}

func TestService(t *testing.T) {
	wrap(func() {
		fmt.Println(service())
		otherTask()
	})()
}

// chan >> channel
func AsyncService() chan string { // 该函数返回一个通道, 消息类型是 string
	retCh := make(chan string) // 声明一个无缓冲 chan通道, 这个通道的类型是 string, 通道用于协程间的通信
	// // ⚠️ 使用无缓冲通道时, 如果协程1 没有立即接受协程 2 通过管道发送的消息，就会阻塞，反之亦然
	//retCh := make(chan string, 1) // 声明一个缓冲容量为 1 的chan通道, 这个通道的类型是 string, 通道用于协程间的通信
	go func() { // 此处执行 匿名协程函数
		ret := service()
		fmt.Println("---AsyncService---returned result:", ret)
		retCh <- ret //通过通道将 service()函数返回的值传递给 通道 retCh
		fmt.Println("---AsyncService---service exited")
	}()
	return retCh // 返回这个通道
}

func TestAsyncService(t *testing.T) {
	wrap(func() {
		retCh := AsyncService()
		otherTask()
		fmt.Println(<-retCh)
	})()
}

func wrap(inner func()) func() {
	return func() {
		time1 := time.Now().UnixNano() / 1e6
		inner()
		time2 := time.Now().UnixNano() / 1e6
		fmt.Printf("Cost %d", time2-time1)
	}
}
