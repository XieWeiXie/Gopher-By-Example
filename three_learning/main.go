package main

import "fmt"

// declare

func example() {
	fmt.Println("Hello world")
}

// arg

func exampleOneArg(arg int) {
	fmt.Println(arg + 1)
}

// arg list

func exampleListArg(args ...int) {
	for index, value := range args {
		fmt.Println(index, value)
	}
}

// one response

func exampleOneResponse(arg int) int {
	return arg*100 + 1
}

// two responses

func exampleTwoResponses(arg int) (int, int) {
	return arg * 10, arg * 100
}

// name response

func exampleNameResponse(arg int) (sum int) {
	sum = arg * 1000
	return
}

// copy value

func exampleValueCopy(arg int) int {
	arg = arg + 1
	return arg
}

// copy address

func exampleValueAddress(arg *int) int {
	*arg = *arg + 1
	return *arg

}

// func as arg

func funcArg(arg int) int {
	return arg * 100
}

func exampleFuncAsArg(arg int, function func(int) int) int {
	return arg + function(arg)

}

// un name func

var NoNameFunc = func(arg int) int { return arg * 1000 }

// main
func main() {
	example()
	exampleOneArg(12)
	exampleListArg(1, 2, 3, 4, 5)
	exampleOneResponse(10)
	exampleTwoResponses(10)
	exampleNameResponse(10)

	var arg int = 10
	exampleValueCopy(arg)
	fmt.Println(arg)
	exampleValueAddress(&arg)
	fmt.Println(arg)

	fmt.Println(exampleFuncAsArg(10, funcArg))
	fmt.Println(NoNameFunc(10))
}
