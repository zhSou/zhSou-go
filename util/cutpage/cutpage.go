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

	if start < 0 {
		start = 0
	}
	if end >= len(slice) {
		end = len(slice)
	}
	return slice[start:end]
}
