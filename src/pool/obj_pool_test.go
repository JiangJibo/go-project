package pool

import (
	"errors"
	"testing"
	"time"
)

type ReusableObj struct {
	id int
}

type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可复用对象
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{i}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func Testt(t *testing.T) {
	pool := NewObjPool(10)
	pool.GetObj(1)

}
