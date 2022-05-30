package cutpage

import "github.com/pkg/errors"

func calcIndex(pageId int, limitSize int) (start, end int) {
	start = (pageId - 1) * limitSize
	end = start + limitSize
	return
}

func CutPage[T any](slice []T, pageId int, limitSize int) ([]T, error) {
	start, end := calcIndex(pageId, limitSize)

	if end < 0 || start >= len(slice) {
		return nil, errors.New("分页错误")
	}
	if start < 0 {
		start = 0
	}
	if end >= len(slice) {
		end = len(slice) - 1
	}
	return slice[start:end], nil
}
