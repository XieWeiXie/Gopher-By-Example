package service

import (
	"errors"
	"fmt"
	"go-example-for-live/ten/infra/golang/execadapter"
)

var (
	ErrGitCloneFail  = errors.New("git clone fail")
	ErrGitCommitFail = errors.New("git commit fail")
	ErrGitStatusFail = errors.New("git status fail")
	ErrGitAddAllFail = errors.New("git add fail")
)

func GitClone(dir string, addr string, path string) (bool, error) {

	// git@github.com:wuxiaoxiaoshen/Resume.git
	var fullPath string
	if addr == "localhost" {
		fullPath = path
	}
	if addr == "github.com" {
		fullPath = path
	}

	var err error
	err = execadapter.Execute(dir, "git", "clone", fullPath)
	if err != nil {
		fmt.Println(ErrGitCloneFail)
		return false, ErrGitCloneFail
	}
	return true, nil
}

func GitCommit(dir string, message string) (bool, error) {

	var err error

	err = execadapter.Execute(dir, "git", "commit", "-m", message)
	if err != nil {
		fmt.Println(ErrGitCommitFail)
		return false, ErrGitCommitFail
	}
	return true, nil
}

func GitStatus(dir string, message string) (bool, error) {

	var err error

	err = execadapter.Execute(dir, "git", "status")
	if err != nil {
		fmt.Println(ErrGitStatusFail)
		return false, ErrGitStatusFail
	}
	return true, nil
}

func GitAddAll(dir string, file string) (bool, error) {

	var err error

	err = execadapter.Execute(dir, "git", "add", "-A", file)
	if err != nil {
		fmt.Println(ErrGitAddAllFail)
		return false, ErrGitAddAllFail
	}
	return true, nil
}
