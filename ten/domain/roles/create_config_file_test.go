package roles

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateConfigDir(t *testing.T) {
	tt := []struct {
		instanceId string
		path       string
	}{
		{
			instanceId: "fuck1",
			path:       "git@github.com:wuxiaoxiaoshen/" + "Resume.git",
		},
		{
			instanceId: "fuck2",
			path:       "git@github.com:wuxiaoxiaoshen/" + "Resume.git",
		},
	}
	for _, t := range tt {
		fmt.Println(CreateConfigDir(t.instanceId, t.path))
		os.RemoveAll("../../ConfigResume/" + t.instanceId)
	}

}
