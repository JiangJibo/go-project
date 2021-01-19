package error

import (
	"errors"
	"testing"
)

// 定义异常
var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThanTwoError = errors.New("n should be not large than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargeThanTwoError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		if err == LargeThanTwoError {
			t.Error(err)
		} else if err == LessThanTwoError {
			t.Error(err)
		}
	} else {
		t.Log(v)
	}
}

func RemoteTagValue(ints []int, tag int) ([]int, int) {
	x := 0
	for idx, v := range ints {
		if v == tag {
			ints[idx] = 0
			ints[x] = 0
		} else {
			ints[x] = v
			if x != idx {
				ints[idx] = 0
			}
			x++
		}
	}
	return ints, x
}

func TestRemoveTagValues(t *testing.T) {
	t.Log(RemoteTagValue([]int{1, 2, 2, 4}, 4))
}
