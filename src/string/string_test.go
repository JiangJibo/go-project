package string

import (
	"testing"
)

func TestStringLen(t *testing.T) {
	s := "中国"
	t.Log(len(s)) // 6
}

func TestStringRune(t *testing.T) {
	s := "中"
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
	//strings.Split()
}
