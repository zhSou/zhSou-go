package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, []int{}, Sum([]int{}, []int{}))
	assert.Equal(t, []int{1}, Sum([]int{1}, []int{}))
	assert.Equal(t, []int{1, 2, 3, 4}, Sum([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3}, Sum([]int{1, 2}, []int{2, 3}))
}
