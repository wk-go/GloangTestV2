package operators

import (
	"strconv"
	"testing"
)

func _bool(b bool, t *testing.T) bool {
	t.Log("bool:", b)
	return b
}

func TestTheANDOperater(t *testing.T) {
	t.Log("test \"&&\" && operator start ...")
	t.Log("--------------------------")
	if _bool(true, t) && _bool(true, t) {
		t.Log("_bool(true, t) && _bool(true, t)")
	}

	t.Log("--------------------------")
	if _bool(false, t) && _bool(true, t) {
		t.Error("_bool(false, t) && _bool(true, t)")
	} else {
		t.Log("_bool(false, t) && _bool(true, t)")
	}

	t.Log("--------------------------")
	a := []int{1, 2, 3}
	if _bool(false, t) && a[len(a)] == 4 {
		t.Error("_bool(false, t) && a[len(a)] == 4")
	} else {
		t.Log("_bool(false, t) && a[len(a)] == 4")
	}

	t.Log("--------------------------")
	if false && a[3] == 4 {
		t.Error("false && a[3] == 4")
	} else {
		t.Log("false && a[3] == 4")
	}

	t.Log("--------------------------")
	if 1 == 2 && a[3] == 4 {
		t.Error("1 == 2 && a[3] == 4")
	} else {
		t.Log("1 == 2 && a[3] == 4")
	}
	t.Log("--------------------------")
	i := 1
	if i == 2 && a[3] == 4 {
		t.Error("i == 2 && a[3] == 4")
	} else {
		t.Log("i == 2 && a[3] == 4")
	}

	t.Log("--------------------------")
	if a[0] == 2 && a[3] == 4 {
		t.Error("a[0] == 2 && a[3] == 4")
	} else {
		t.Log("a[0] == 2 && a[3] == 4")
	}

	t.Log("--------------------------")
	b := []int{1, 2, 3}
	if b[0] != 1 && a[3] == 4 {
		t.Error("b[0] != 1 && a[3] == 4")
	} else {
		t.Log("b[0] != 1 && a[3] == 4")
	}

	t.Log("--------------------------")

	t.Log("test \"&&\" && operator end!!")
}

type A struct {
	a int
	b int
}

func (a A) String() string {
	return "[A] a:" + strconv.Itoa(a.a) + " b:" + strconv.Itoa(a.b)
}

type B struct {
	a int
	b int
}

func (b B) String() string {
	return "[B] a:" + strconv.Itoa(b.a) + " b:" + strconv.Itoa(b.b)
}

// C 继承A和B 更像A和B
type C struct {
	A
	B
}

// D 继承A和B
type D struct {
	A
	B
}

// D 拥有自己的小性格，覆盖A和B的String方法
func (d D) String() string {
	return "[D] a:" + strconv.Itoa(d.A.a) + " b:" + strconv.Itoa(d.B.b)
}

// 测试结构体多重继承 访问方法和属性的效果
func TestStructMultipleInheritance(t *testing.T) {
	c := C{A: A{a: 1, b: 2}, B: B{a: 3, b: 4}}
	t.Log(c)            // 分别调用A和B的String方法
	t.Log(c.A.String()) // 调用A的String方法

	d := D{A: A{a: 11, b: 12}, B: B{a: 13, b: 14}}
	t.Log(d)          // 调用D的String方法
	t.Log(d.String()) // 调用D的String方法
}
