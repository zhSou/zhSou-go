package cutpage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCutPage(t *testing.T) {
	s1 := []int{5, 2, 8, 3, 1, 2}
	// 按长度为2取第1页，第2页
	assert.Equal(t, []int{5, 2}, CutPage[int](s1, 1, 2))
	assert.Equal(t, []int{8, 3}, CutPage[int](s1, 2, 2))
	// 按长度为3取第1,2,3页
	assert.Equal(t, []int{5, 2, 8}, CutPage[int](s1, 1, 3))
	assert.Equal(t, []int{3, 1, 2}, CutPage[int](s1, 2, 3))
	assert.Equal(t, []int{}, CutPage[int](s1, 3, 3))
	// 按长度为4取第1,2,3页
	assert.Equal(t, []int{5, 2, 8, 3}, CutPage[int](s1, 1, 4))
	assert.Equal(t, []int{1, 2}, CutPage[int](s1, 2, 4))
	assert.Equal(t, []int{}, CutPage[int](s1, 3, 4))

	// 按长度为0取第-1页，第0页，第1页
	assert.Equal(t, []int{}, CutPage[int](s1, -1, 0))
	assert.Equal(t, []int{}, CutPage[int](s1, 0, 0))
	assert.Equal(t, []int{}, CutPage[int](s1, 1, 0))
	// 按长度为-1取第-1页，第0页，第1页
	assert.Equal(t, []int{}, CutPage[int](s1, -1, -1))
	assert.Equal(t, []int{}, CutPage[int](s1, 0, -1))
	assert.Equal(t, []int{}, CutPage[int](s1, 1, -1))

	// 空数据集
	s2 := make([]int, 0)
	assert.Equal(t, []int{}, CutPage[int](s2, 1, 2))
	assert.Equal(t, []int{}, CutPage[int](s2, 2, 1))
}
