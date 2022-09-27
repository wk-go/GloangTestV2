package unsafe

import "testing"

var (
	s = []string{
		"Hello world",
		"ABCDEFG",
	}
)

func TestString2bytes(t *testing.T) {
	for _, v := range s {
		b := []byte(v)
		for i, _byte := range String2bytes(v) {
			if b[i] != _byte {
				t.Errorf(v, b, String2bytes(v))
			}
		}

	}
}

func TestBytes2String(t *testing.T) {
	for _, v := range s {
		b := []byte(v)
		if v != Bytes2String(b) {
			t.Errorf(v, b, Bytes2String(b))
		}
	}
}
