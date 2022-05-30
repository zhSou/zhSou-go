package filesystem

import (
	"log"
	"os"
	"strings"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDir(filepath string) error {
	i := strings.LastIndexFunc(filepath, func(r rune) bool {
		return r == '\\'
	})
	dirPath := filepath[:i]
	exist, err := Exists(dirPath)
	if err != nil {
		return err
	}
	if exist {
		log.Println("路径已存在：", dirPath)
		return nil
	}
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}
	log.Println("新建文件夹：", dirPath)
	return nil
}

func DeleteFile(filepath string) bool {
	err := os.Remove(filepath)

	if err == nil {
		log.Println("成功删除文件：", filepath)
		return true
	}
	return false
}
