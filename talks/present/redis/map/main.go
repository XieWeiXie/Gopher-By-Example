package main

import "fmt"

func main() {
	var m = make(map[int]string)
	m = map[int]string{
		1: "x",
		2: "y",
		3: "z",
	}
	for key, val := range m {
		fmt.Println(key, val)
	}
	if val, ok := m[2]; ok {
		fmt.Println(val)
	}
}
