package ioutiladapter

import (
	"fmt"
	"testing"
)

func TestReadAllCfgFromFile(t *testing.T) {

	tt := []struct {
		fileName string
	}{
		{
			fileName: "xiewei.txt",
		},
		{
			fileName: "test.txt",
		},
	}

	for _, t := range tt {
		var (
			out []byte
			err error
		)
		out, err = ReadAllCfgFromFile(t.fileName)
		fmt.Println(string(out), err)
	}

}

func TestWriteCfgIntoFile(t *testing.T) {
	tt := []struct {
		data     []byte
		fileName string
	}{
		{
			data: []byte(`func main() {
    var b StructB

    err := yaml.Unmarshal([]byte(data), &b)
    if err != nil {
        log.Fatalf("cannot unmarshal data: %v", err)
    }
    fmt.Println(b.A)
    fmt.Println(b.B)
}`),
			fileName: "test.txt",
		},
		{
			data:     []byte(`Thank you 2`),
			fileName: "xiewei1.txt",
		},
	}
	for _, t := range tt {
		fmt.Println(WriteCfgIntoFile(t.fileName, t.data))
	}

}
