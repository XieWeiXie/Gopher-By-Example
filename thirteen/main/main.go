package main

import (
	"flag"
	"fmt"
	"go-example-for-live/thirteen/infra"
	"strings"
)

func Pinyin() {

	var name string

	flag.StringVar(&name, "n", "谢伟", "")
	flag.Parse()
	pys := []string{}

	for _, arg := range flag.Args() {
		for _, v := range arg {
			py, ok := infra.PinyinDict[v]
			if ok {
				pyList := strings.Split(py, ",")
				if len(pyList) > 0 {
					pys = append(pys, pyList[0])
				}
			}
		}

	}

	fmt.Println(pys)

}

func main() {
	var dictMap = map[int32]string{}

	dictMap[0x8c22] = "xiè"
	dictMap[0x4f1f] = "wěi"

	var a = "谢伟"
	pys := [][]string{}
	for _, v := range a {
		py, ok := dictMap[v]
		pyr := []string{}
		if ok {
			pyr = strings.Split(py, ",")
		}
		pys = append(pys, pyr)
	}

	//fmt.Println(pys)
	Pinyin()

}
