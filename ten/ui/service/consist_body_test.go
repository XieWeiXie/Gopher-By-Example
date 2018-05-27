package service

import (
	"fmt"
	"testing"
)

func TestConsistBody(t *testing.T) {
	tt := []struct {
		key      string
		fileName string
	}{
		{
			key:      "json",
			fileName: "../../config/cfg.yml",
		},
		{
			key:      "yaml",
			fileName: "../../config/cfg.yml",
		},
		{
			key:      "object",
			fileName: "../../config/cfg.yml",
		},
		{
			key:      "paragraph",
			fileName: "../../config/cfg.yml",
		},
	}

	for _, t := range tt {
		fmt.Println(ConsistBody(t.key, t.fileName))
	}
}
