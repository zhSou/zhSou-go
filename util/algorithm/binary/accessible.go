package binary

// RandomAccessible 可随机访问的数据接口
type RandomAccessible[T any] interface {
	Get(id int) T
	Len() int
}

type SliceAccessible[T any] struct {
	src []T
}

func NewSliceAccessible[T any](src []T) *SliceAccessible[T] {
	return &SliceAccessible[T]{src}
}

func (s *SliceAccessible[T]) Get(id int) T {
	return s.src[id]
}

func (s *SliceAccessible[T]) Len() int {
	return len(s.src)
}
