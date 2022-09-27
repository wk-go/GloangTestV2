package unsafe

import "unsafe"

func String2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
