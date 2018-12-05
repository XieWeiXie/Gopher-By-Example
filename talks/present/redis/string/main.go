package main

import "fmt"

func main() {
	var (
		name   string
		number int
	)
	name = "xieWei"
	number = 12
	fmt.Println(name, name[1:], len(name))
	fmt.Println(number, number+1, float32(number)+10.2)
	fmt.Println(number, number-1, float32(number)-10.2)
}
