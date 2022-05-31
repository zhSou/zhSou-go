package cutpage

func calcIndex(pageId int, limitSize int) (start, end int) {
	start = (pageId - 1) * limitSize
	end = start + limitSize
	return
}

func CutPage[T any](slice []T, pageId int, limitSize int) []T {
	if len(slice) == 0 {
		return slice
	}
	if limitSize <= 0 || pageId <= 0 {
		return []T{}
	}
	start, end := calcIndex(pageId, limitSize)

	// 起始地址越界
	if start >= len(slice) {
		return []T{}
	}
	if end >= len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}
