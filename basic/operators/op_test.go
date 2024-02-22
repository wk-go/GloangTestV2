package operators

import "testing"

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
