package binary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceAccessible_Get(t *testing.T) {
	sa := SliceAccessible[int]{[]int{1, 2, 3, 4}}
	assert.Equal(t, 1, sa.Get(0))
	assert.Equal(t, 4, sa.Get(3))
}

func TestSliceAccessible_Len(t *testing.T) {
	sa := SliceAccessible[int]{[]int{1, 2, 3, 4}}
	assert.Equal(t, 4, sa.Len())

	sa = SliceAccessible[int]{[]int{}}
	assert.Equal(t, 0, sa.Len())
}
