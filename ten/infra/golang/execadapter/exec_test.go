package execadapter

import "testing"

func TestExecute(t *testing.T) {
	tt := []struct {
		dir  string
		cmd  string
		args string
	}{
		{
			dir:  ".",
			cmd:  "git",
			args: "status",
		},
		{
			dir:  ".",
			cmd:  "python",
			args: "--version",
		},
		{
			dir:  ".",
			cmd:  "git",
			args: "branch",
		},
	}

	for _, t := range tt {
		Execute(t.dir, t.cmd, t.args)
	}
}
