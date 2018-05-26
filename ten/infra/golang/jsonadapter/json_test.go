package jsonadapter

import (
	"fmt"
	"testing"
)

type Value struct {
	Name  string
	Order string
}

func TestUnmarshal(t *testing.T) {

	tt := []struct {
		data  []byte
		value interface{}
		err   error
	}{
		{
			data:  []byte(`{"Name": "Platypus", "Order": "Monotremata"}`),
			value: &Value{},
			err:   nil,
		},
		{
			data:  []byte(`{"Name": "Quoll",    "Order": "Dasyuromorphia"}`),
			value: &Value{},
			err:   nil,
		},
	}
	for _, t := range tt {
		if err := Unmarshal(t.data, t.value); err != t.err {
			fmt.Println("Wrong")
		}
		fmt.Println(t.value)
	}
}

func TestMarshal(t *testing.T) {
	tt := []struct {
		v   Value
		err error
	}{
		{
			v:   Value{"Platypus", "Monotremata"},
			err: nil,
		},
		{
			v:   Value{"Quoll", "Dasyuromorphia"},
			err: nil,
		},
	}
	for _, t := range tt {
		data, err := Marshal(t.v)
		fmt.Println(string(data), err)
		if err != t.err {
			fmt.Println("Wrong")
		}
	}
}
