package groutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGrouting(t *testing.T) {
	for i := 0; i < 10; i++ {
		/*go func(i int) {
			fmt.Println(i)
		}(i)*/
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}

func TestCounter(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}
