package yamladapter

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	tt := []struct {
		Name   string `yaml:"name"`
		Age    string `yaml:"age"`
		School string `yaml:"school"`
	}{
		{
			Name:   "xieWei",
			Age:    "12",
			School: "shangHaiUniversity",
		},
		{
			Name:   "xieWei2",
			Age:    "20",
			School: "pekingUniversity",
		},
	}

	for _, t := range tt {
		var (
			out []byte
			err error
		)
		out, err = Marshal(t)
		if err == nil {
			fmt.Println(string(out))
		}
	}

}

type Infomation struct {
	Name   string `yaml:"name"`
	Age    string `yaml:"age"`
	School string `yaml:"school"`
}

func TestUnmarshal(t *testing.T) {

	tt := []struct {
		Contents []byte
		Info     Infomation
	}{
		{
			Contents: []byte(`name: xieWei
age: "12"
school: shangHaiUniversity`),
			Info: Infomation{},
		},
		{
			Contents: []byte(`name: xieWei2
age: "20"
school: pekingUniversity`),
			Info: Infomation{},
		},
	}
	for _, t := range tt {
		Unmarshal(t.Contents, &t.Info)
		fmt.Println(t.Info.Name)
	}

}
