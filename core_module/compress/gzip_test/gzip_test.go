package gzip_test

import (
	"compress/gzip"
	"encoding/json"
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

type Data struct {
	Name  string            `json:"name"`
	Data  string            `json:"data"`
	Files map[string][]byte `json:"files"`
}

func TestCompressStructData(t *testing.T) {
	data := Data{
		Name: "ftest",
		Data: "filename:gzip_test.go",
		Files: map[string][]byte{
			"gzip_test.go": nil,
		},
	}
	t.Log(data)
	t.Log(os.Getwd())

	fileContent, err := os.ReadFile("./gzip_test.go")
	if err != nil {
		t.Error(err)
	}
	data.Files["gzip_test.go"] = fileContent
	_json, _ := json.Marshal(data)
	gzipFile, err := os.Create("test/data.gz")
	defer gzipFile.Close()
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()
	gzipWriter.Write(_json)

}

func TestUnCompressStructData(t *testing.T) {
	gzipFile, err := os.Open("./test/data.gz")
	if err != nil {
		t.Error(err)
	}
	defer gzipFile.Close()
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		t.Error(err)
	}
	defer gzipReader.Close()
	_data, err := io.ReadAll(gzipReader)
	if err != nil {
		t.Error(err)
	}
	data := &Data{}
	if err := json.Unmarshal(_data, data); err != nil {
		t.Error(err)
	}

	t.Log(data)

	outfileWriter, err := os.Create("./test/_data.json")
	if err != nil {
		t.Error(err)
	}
	defer outfileWriter.Close()
	outfileWriter.Write(_data)
}
