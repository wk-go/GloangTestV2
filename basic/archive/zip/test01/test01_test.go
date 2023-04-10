package main

import (
	"bytes"
	"os"
	"testing"
)

func TestZipToWriter(t *testing.T) {
	var (
		dest = "myFiles1.zip"
		buf  = new(bytes.Buffer)
	)

	ZipToWriter(buf, "zipdir")
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	var b2 = buf.Bytes()
	t.Log(b2)
	if _, err := buf.WriteTo(f); err != nil {
		t.Error(err)
	}
	t.Log("complete")
}

func TestZip(t *testing.T) {
	dest := "myFiles.zip"
	err := Zip(dest, "zipdir")
	if err != nil {
		t.Fatal(err)
	}
}
func TestUnZip(t *testing.T) {
	err := UnZip("myFiles.zip", "unzip")
	if err != nil {
		t.Fatal(err)
	}
}
