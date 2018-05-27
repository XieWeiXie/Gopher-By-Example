package service

import (
	"fmt"
	"go-example-for-live/ten/infra/golang/osadapater"
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	tt := []struct {
		dir  string
		addr string
		path string
	}{
		{
			dir:  ".",
			addr: "localhost",
			path: "../../../../Resume/.git",
		},
		{
			dir:  "./xiewei",
			addr: "github.com",
			path: "Resume.git",
		},
	}
	osadapater.RDir("xiewei")
	osadapater.RDir("Resume")
	err := os.Mkdir("xiewei", 0644)
	if err != nil {
		fmt.Println("can not create dir")
		return
	}
	for _, t := range tt {

		fmt.Println(GitClone(t.dir, t.addr, t.path))
	}
	osadapater.RDir("./xiewei")
	osadapater.RDir("./Resume")

}
