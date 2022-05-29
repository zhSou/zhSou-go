package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstBigger(t *testing.T) {
	s1 := []int{1, 6, 21, 90, 100}
	// 正常情况
	assert.Equal(t, FindFirstBigger(&SliceAccessible{s1}, 1), 1)
	assert.Equal(t, FindFirstBigger(&SliceAccessible{s1}, 21), 3)
	// 比4大的数找不到
	assert.Equal(t, FindFirstBigger(&SliceAccessible{s1}, 100), -1)

	// 不存在任何数字，故查找到的索引值恒为-1
	assert.Equal(t, FindFirstBigger(&SliceAccessible{[]int{}}, 1), -1)
	assert.Equal(t, FindFirstBigger(&SliceAccessible{[]int{}}, 4), -1)
}
