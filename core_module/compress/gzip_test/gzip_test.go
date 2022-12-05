package gzip_test

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	outputFile, _ := os.Create("./test/test.gz")
	defer outputFile.Close()
	gzipWriter := gzip.NewWriter(outputFile) //gzipWriter:需要操作的句柄
	defer gzipWriter.Close()
	//写入gizp writer数据时，它会依次压缩数据并写入到底层的文件中
	gzipWriter.Write([]byte("hello world!\n"))
	log.Println("success")
}

func TestUncomporess(t *testing.T) {
	gzipFile, _ := os.Open("./test/test.gz")
	defer gzipFile.Close()
	gzipReader, _ := gzip.NewReader(gzipFile)
	defer gzipReader.Close()
	outfileWriter, err := os.Create("./test/unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outfileWriter.Close()
	// 复制内容
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
