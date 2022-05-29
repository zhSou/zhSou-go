package binary

// RandomAccessible 可随机访问的数据接口
type RandomAccessible interface {
	Get(id int) int
	Len() int
}

type SliceAccessible struct {
	src []int
}

func (s *SliceAccessible) Get(id int) int {
	return s.src[id]
}

func (s *SliceAccessible) Len() int {
	return len(s.src)
}
