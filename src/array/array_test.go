package array

import (
	"fmt"
	"testing"
)

func TestArraySliceGrowing(t *testing.T) {
	x := [3]int{1, 2, 3}
	t.Log(x)
	s := []int{}
	fmt.Print(len(s), cap(s))
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
	summer = append(summer, "Xxx")
	t.Log(summer, len(summer), cap(summer))

	t.Log(len(year), cap(year))
	year = append(year, "yyyy")
	t.Log(len(year), cap(year))
}

func TestCompareSlice(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	if a == b {

	}
}
