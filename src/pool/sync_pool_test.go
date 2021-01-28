package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println("v:", v)
	pool.Put(3)
	// 官方Go runtime 1.13将对sync.Pool中的对象回收时机策略做出调整。
	//在1.12版本及以前的版本中，在每轮垃圾回收过程中，每个sync.Pool实例中的所有缓存对象都将被无条件回收掉。
	//从1.13版本开始，如果一个sync.Pool实例在上一轮垃圾回收过程结束之后仍然被使用过，则其中的缓存对象将不会被回收掉。
	//此举对于使用sync.Pool来提升效率的程序来说，将大大减少周期性的因为缓存被清除而造成的瞬时效率下降。
	runtime.GC() // GC会清除sync.Pool 中缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println("v1:", v1)
	v2, _ := pool.Get().(int)
	fmt.Println("v2", v2)
}

func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 10
		},
	}

	pool.Put(1)
	pool.Put(2)
	pool.Put(3)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
