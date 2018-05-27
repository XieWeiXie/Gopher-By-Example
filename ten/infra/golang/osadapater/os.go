package osadapater

import (
	"errors"
	"fmt"
	"os"
)

// 功能点: 创建文件夹, 判断文件是否存在

var (
	ErrFileInfo   = errors.New("file stat error")
	ErrDirNoExist = errors.New("dir not exist ")
	ErrDirExist   = errors.New("dir exist")
	ErrFileExist  = errors.New("file exist")
)

func MKdir(name string) error {

	var (
		ok  bool
		err error
	)

	ok, err = FExist(name)
	if err != nil || ok == false {
		return os.Mkdir(name, 0650)
	} else {
		return ErrDirExist
	}

}

func RDir(path string) error {

	var ok bool

	ok = DExist(path)
	if ok == true {
		return os.RemoveAll(path)
	} else {
		return ErrDirNoExist
	}

}

func FExist(path string) (bool, error) {

	var (
		err error
	)
	_, err = os.Stat(path)
	if os.IsNotExist(err) && err != nil {
		return false, ErrFileInfo
	}

	return true, nil
}

func DExist(path string) bool {
	var (
		fileInfo os.FileInfo
		err      error
	)
	fileInfo, err = os.Stat(path)
	if err != nil {
		fmt.Println(ErrDirNoExist)
		return false
	}
	return fileInfo.IsDir()
}

func CFile(name string) error {

	var (
		ok  bool
		err error
	)

	ok, err = FExist(name)
	if ok != true || err != nil {
		_, err = os.Create(name)
		return err
	} else {
		return ErrFileExist
	}

}
