package main

import (
	"fmt"
	"io/ioutil"
)

// err

func example(arg int) (int, error) {
	return arg + 1, nil
}

// errors.New
func exampleReadAllError() error {
	_, err := ioutil.ReadAll(nil)
	return err
}

// fmt.Errorf

func exampleErrorf(arg int) error {
	return fmt.Errorf("%d", arg)

}

func examplePanic() {
	panic("123")
}

func examplePanicOther() {
	panic("1234")
}

func main() {
	if key, err := example(2); err != nil {
		fmt.Println(key)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("recieve error: %v", err))
		}
	}()
	examplePanic()
	examplePanicOther()
	examplePanic()

}
