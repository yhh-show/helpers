package file

import (
	"errors"
	"os"
	"path/filepath"
)

func Exists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Find(name string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for i := 0; i < 100; i++ {
		dir := pwd + "/" + name
		// 判断文件是否存在
		exists, err := Exists(dir)
		if err != nil {
			return "", err
		}
		if !exists {
			// 到顶了
			if pwd == "/" {
				break
			}
			pwd, err = filepath.Abs(pwd + "/..")
			if err != nil {
				return "", err
			}
			if pwd == "/" {
				pwd = ""
			}
			continue
		}
		return dir, nil
	}
	return "", errors.New(name + " not found")
}
