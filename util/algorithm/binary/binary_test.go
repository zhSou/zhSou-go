package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstBigger(t *testing.T) {
	s1 := []int{1, 6, 21, 90, 100}
	as := &SliceAccessible[int]{s1}
	// 正常情况
	assert.Equal(t, 0, FindFirstBigger[int](as, -1))
	assert.Equal(t, 0, FindFirstBigger[int](as, 0))

	assert.Equal(t, 1, FindFirstBigger[int](as, 1))
	assert.Equal(t, 1, FindFirstBigger[int](as, 4))
	assert.Equal(t, 3, FindFirstBigger[int](as, 21))

	assert.Equal(t, 4, FindFirstBigger[int](as, 91))
	// 比100,104大的数找不到，返回-1
	assert.Equal(t, -1, FindFirstBigger[int](as, 100))
	assert.Equal(t, -1, FindFirstBigger[int](as, 104))

	// 不存在任何数字，故查找到的索引值恒为-1
	assert.Equal(t, -1, FindFirstBigger[int](NewSliceAccessible([]int{}), 1))
	assert.Equal(t, -1, FindFirstBigger[int](NewSliceAccessible([]int{}), 4))
}
