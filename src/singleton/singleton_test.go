package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type SingleTon struct {
}

var singletonInstance *SingleTon
var once sync.Once

func GetSingleTonObj() *SingleTon {
	once.Do(func() {
		fmt.Println("Create Obj")
		singletonInstance = new(SingleTon)
	})
	return singletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingleTonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
