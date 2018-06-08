package engine

import (
	"go-example-for-live/fifteen/parse/v2ex"
	"testing"
)

func TestRun(test *testing.T) {
	request := Request{
		Url:       "https://www.v2ex.com/?tab=jobs",
		ParseFunc: v2ex.JobParse,
	}
	tt := []struct {
		seed Request
	}{
		{
			seed: request,
		},
	}

	for _, t := range tt {
		Run(t.seed)
	}
}
