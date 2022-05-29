package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstBigger(t *testing.T) {
	s1 := []int{1, 6, 21, 90, 100}
	// 正常情况
	assert.Equal(t, 0, FindFirstBigger(&SliceAccessible{s1}, -1))
	assert.Equal(t, 0, FindFirstBigger(&SliceAccessible{s1}, 0))

	assert.Equal(t, 1, FindFirstBigger(&SliceAccessible{s1}, 1))
	assert.Equal(t, 1, FindFirstBigger(&SliceAccessible{s1}, 4))
	assert.Equal(t, 3, FindFirstBigger(&SliceAccessible{s1}, 21))

	assert.Equal(t, 4, FindFirstBigger(&SliceAccessible{s1}, 91))
	// 比100,104大的数找不到，返回-1
	assert.Equal(t, -1, FindFirstBigger(&SliceAccessible{s1}, 100))
	assert.Equal(t, -1, FindFirstBigger(&SliceAccessible{s1}, 104))

	// 不存在任何数字，故查找到的索引值恒为-1
	assert.Equal(t, -1, FindFirstBigger(&SliceAccessible{[]int{}}, 1))
	assert.Equal(t, -1, FindFirstBigger(&SliceAccessible{[]int{}}, 4))
}
