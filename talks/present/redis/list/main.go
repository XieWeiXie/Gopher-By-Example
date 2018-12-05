package main

import "fmt"

func main() {
	var l = [4]int{1, 2, 3, 4}
	fmt.Println(l[2], len(l), l[2:])
	newSlice := l[2:]
	newSlice = append(newSlice, 5)
	fmt.Println(newSlice[2], len(newSlice), newSlice)
}
