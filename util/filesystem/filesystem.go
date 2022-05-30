package filesystem

import (
	"os"
	"strings"
)

func MakeDir(filepath string) error {
	i := strings.LastIndexFunc(filepath, func(r rune) bool {
		return r == '\\'
	})
	dirPath := filepath[:i]
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}
	return nil
}
