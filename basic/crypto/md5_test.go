package crypto

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5Sum(t *testing.T) {
	data := []string{
		"Hello world",
		"你好，世界！",
	}
	for _, v := range data {
		hash := md5.Sum([]byte(v))
		d := md5.New()
		d.Write([]byte(v))
		hash2 := d.Sum(nil)
		fmt.Printf("%x:%x\n", hash, hash2)
	}
}
