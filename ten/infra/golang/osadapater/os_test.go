package osadapater

import (
	"fmt"
	"testing"
)

func TestMKdir(t *testing.T) {
	tt := []struct {
		path string
	}{
		{
			path: "./name",
		},
		{
			path: "./xie",
		},
	}
	for _, t := range tt {
		err := MKdir(t.path)
		if err != nil {
			fmt.Println("wrong", err)
		}
	}
}

func TestRDir(t *testing.T) {
	tt := []struct {
		path string
	}{
		{
			path: "./name",
		},
		{
			path: "./xie",
		},
	}
	for _, t := range tt {
		RDir(t.path)
	}
}

func TestFExist(t *testing.T) {
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
		fmt.Println(FExist(t.fileName))
	}

}

func TestDExist(t *testing.T) {

	tt := []struct {
		dirName string
	}{
		{
			dirName: "./xiewei",
		},
		{
			dirName: "./test",
		},
	}
	for _, t := range tt {
		fmt.Println(DExist(t.dirName))
	}

}

func TestCFile(t *testing.T) {
	tt := []struct {
		fileName string
	}{
		{
			fileName: "./xiewei.txt",
		},
		{
			fileName: "./test.txt",
		},
	}
	for _, t := range tt {
		fmt.Println(CFile(t.fileName))
	}

}
