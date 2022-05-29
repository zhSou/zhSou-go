package binary

// FindFirstBigger 在一个有序表中查找比target大的最小值。若找到则返回非负整数的索引值，否则返回-1
func FindFirstBigger(data RandomAccessible, target int) (id int) {
	if data.Len() == 0 {
		return -1
	}
	left := 0
	right := data.Len() - 1
	for left < right {
		mid := (right-left)/2 + left
		if data.Get(mid) > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if data.Get(left) <= target {
		return -1
	}
	return left
}
