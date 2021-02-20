package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int64, reflect.Int32, reflect.Int:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown")
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)
}
