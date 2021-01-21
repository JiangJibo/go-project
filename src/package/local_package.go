package _package

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("my_series init")
}

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

func Square(n int) int {
	return n * n
}
