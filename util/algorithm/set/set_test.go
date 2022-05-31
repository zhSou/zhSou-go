package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCross(t *testing.T) {
	assert.Equal(t, []int{}, Cross([]int{}, []int{}))
	assert.Equal(t, []int{}, Cross([]int{1}, []int{}))
	assert.Equal(t, []int{}, Cross([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{2, 3}, Cross([]int{1, 2, 3, 4}, []int{2, 3, 8, 7}))
}

func TestSum(t *testing.T) {
	assert.Equal(t, []int{}, Sum([]int{}, []int{}))
	assert.Equal(t, []int{1}, Sum([]int{1}, []int{}))
	assert.Equal(t, []int{1, 2, 3, 4}, Sum([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3}, Sum([]int{1, 2}, []int{2, 3}))
}

func TestExclude(t *testing.T) {
	assert.Equal(t, []int{}, Exclude([]int{1, 2}, []int{1, 2}))
	assert.Equal(t, []int{}, Exclude([]int{}, []int{}))
	assert.Equal(t, []int{1, 2, 3, 4}, Exclude([]int{1, 2, 3, 4}, []int{}))
	assert.Equal(t, []int{1, 2}, Exclude([]int{1, 2, 3, 4}, []int{3, 4}))
	assert.Equal(t, []int{1, 4}, Exclude([]int{1, 2, 3, 4}, []int{2, 3}))
}
