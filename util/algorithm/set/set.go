package set

// Deduplication 删除切片中的重复元素(不保证顺序)
func Deduplication[T comparable](s []T) []T {
	mp := map[T]struct{}{}
	for _, item := range s {
		mp[item] = struct{}{}
	}
	var ret []T
	for item, _ := range mp {
		ret = append(ret, item)
	}
	return ret
}

/// 有序表的集合运算

func Cross(s1, s2 []int) []int {
	var res []int
	res = []int{}
	p1, p2 := 0, 0
	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1] < s2[p2] {
			p1++
		} else if s1[p1] == s2[p2] {
			res = append(res, s1[p1])
			p1++
			p2++
		} else {
			p2++
		}
	}
	return res
}

func Sum(s1, s2 []int) []int {
	var res []int
	res = []int{}
	p1, p2 := 0, 0
	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1] < s2[p2] {
			res = append(res, s1[p1])
			p1++
		} else if s1[p1] == s2[p2] {
			res = append(res, s1[p1])
			p1++
			p2++
		} else {
			res = append(res, s2[p2])
			p2++
		}
	}
	for p1 < len(s1) {
		res = append(res, s1[p1])
		p1++
	}
	for p2 < len(s2) {
		res = append(res, s2[p2])
		p2++
	}
	return res
}

func Exclude(source, exclude []int) []int {
	var res []int
	res = []int{}
	p1, p2 := 0, 0
	for p1 < len(source) && p2 < len(exclude) {
		if source[p1] < exclude[p2] {
			res = append(res, source[p1])
			p1++
		} else if source[p1] == exclude[p2] {
			p1++
			p2++
		} else {
			p2++
		}
	}
	for p1 < len(source) {
		res = append(res, source[p1])
		p1++
	}
	return res
}
