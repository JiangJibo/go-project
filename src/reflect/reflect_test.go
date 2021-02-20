package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type MyData struct {
	aByte   byte
	aShort  int16
	anInt32 int32
	aSlice  []byte
}

//内存对齐是一种牺牲空间节约时间的做法，但是我们一定要注意不要浪费空间，
//聚合类型定义的时候一定要将占用内从空间小的类型放在前面。
func TestReflect(t *testing.T) {
	typ := reflect.TypeOf(MyData{})
	fmt.Printf("Struct is %d bytes long \n", typ.Size())
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n", field.Name, field.Offset, field.Type.Size(), field.Type.Align())
	}

	data := MyData{
		aByte:   0x1,
		aShort:  0x0203,
		anInt32: 0x04050607,
		aSlice: []byte{
			0x08, 0x09, 0x0a,
		},
	}
	dataBytes := (*[32]byte)(unsafe.Pointer(&data))
	fmt.Printf("Bytes are %#v\n", dataBytes)

	dataSlice := *(*reflect.SliceHeader)(unsafe.Pointer(&data.aSlice))
	fmt.Printf("Slice data is %#v\n", (*[3]byte)(unsafe.Pointer(dataSlice.Data)))

	var a int = 1
	fmt.Println(&a)
	b := 2
	fmt.Print(&b)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2 ?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3 ?", reflect.DeepEqual(s1, s3))

}
