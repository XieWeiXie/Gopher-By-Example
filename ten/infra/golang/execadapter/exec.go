package execadapter

import (
	"errors"
	"fmt"
	"os/exec"
)

// 功能点：任意目录下执行终端命令

var (
	ErrCmdNotFound = errors.New("cmd not found")
	ErrCmdFail     = errors.New("cmd exec fail")
)

func Execute(dir string, cmd string, args ...string) error {

	var (
		ok  string
		err error
	)

	ok, err = exec.LookPath(cmd)
	if err != nil {
		fmt.Println(ErrCmdNotFound, ok)
		return ErrCmdNotFound
	}

	var executeCmd *exec.Cmd
	executeCmd = exec.Command(cmd, args...)
	executeCmd.Dir = dir

	var outString []byte

	outString, err = executeCmd.CombinedOutput()
	if err != nil {
		fmt.Println(ErrCmdFail)
		return ErrCmdFail
	}

	fmt.Println(string(outString))
	return nil

}
