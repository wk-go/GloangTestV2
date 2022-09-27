package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Programmer struct {
	name     string
	language string
}

type Programmer2 struct {
	name     string
	age      int
	language string
}

func TestChangeField1(t *testing.T) {
	p := Programmer{"stefno", "go"}
	fmt.Println(p)
	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)
}

func TestChangeField2(t *testing.T) {
	p := Programmer2{"stefno", 18, "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "qcrao"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"

	fmt.Println(p)

	//如果不在同一个包里面可以通过unsafe.Sizeof根据成员类型获取成员大小
	lang = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "Java"
	fmt.Println(p)
}

func TestGetSliceInfo(t *testing.T) {
	var x *int
	v := reflect.TypeOf(x)
	fmt.Println(v.Kind(), v.Size())
	fmt.Println(v.Elem().Kind(), v.Elem().Size())
	//测试获取slice的长度和容量
	s := make([]int, 9, 20)
	Len := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s))

	Cap := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s))
}

func TestGetMapInfo(t *testing.T) {
	_map := make(map[string]int)
	_map["qcrao"] = 100
	_map["stefno"] = 18

	//因为map本身就是指针所以使用的时候就变成二级指针了。
	count := **(**int)(unsafe.Pointer(&_map))
	fmt.Println(count, len(_map))
}
