package download

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{
			"https://gocn.io/",
			nil,
		},
		{
			"https://www.jianshu.com/u/58f0817209aa",
			nil,
		},
	}

	for _, test := range tests {
		_, err := Download(test.url)
		if err != test.err {
			fmt.Println("error")
		}

	}
}
