package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	crossTestCases = []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{
			input1: []int{},
			input2: []int{},
			output: []int{},
		},
		{
			input1: []int{1},
			input2: []int{},
			output: []int{},
		},
		{
			input1: []int{},
			input2: []int{1},
			output: []int{},
		},

		{
			input1: []int{1, 2},
			input2: []int{1, 3},
			output: []int{1},
		},
		{
			input1: []int{1, 2},
			input2: []int{3, 4, 5},
			output: []int{},
		},
		{
			input1: []int{1, 2, 3, 4},
			input2: []int{2, 3, 7, 8},
			output: []int{2, 3},
		},
	}

	sumTestCases = []struct {
		input1 []int
		input2 []int
		output []int
	}{
		{
			input1: []int{},
			input2: []int{},
			output: []int{},
		},
		{
			input1: []int{1},
			input2: []int{},
			output: []int{1},
		},
		{
			input1: []int{},
			input2: []int{1},
			output: []int{1},
		},

		{
			input1: []int{1, 2},
			input2: []int{1, 3},
			output: []int{1, 2, 3},
		},
		{
			input1: []int{1, 2},
			input2: []int{3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input1: []int{1, 2, 3, 4},
			input2: []int{2, 3, 7, 8},
			output: []int{1, 2, 3, 4, 7, 8},
		},
	}

	excludeTestCases = []struct {
		source  []int
		exclude []int
		output  []int
	}{
		{
			source:  []int{1, 2},
			exclude: []int{1, 2},
			output:  []int{},
		},
		{
			source:  []int{},
			exclude: []int{1},
			output:  []int{},
		},
		{
			source:  []int{1, 2, 3, 4},
			exclude: []int{1, 3},
			output:  []int{2, 4},
		},
		{
			source:  []int{1, 2},
			exclude: []int{1, 2, 3},
			output:  []int{},
		},
	}
)

func TestCross(t *testing.T) {
	for _, testCase := range crossTestCases {
		assert.Equal(t, testCase.output, Cross(testCase.input1, testCase.input2), testCase)
	}
}

func TestSum(t *testing.T) {
	for _, testCase := range sumTestCases {
		assert.Equal(t, testCase.output, Sum(testCase.input1, testCase.input2), testCase)
	}
}

func TestExclude(t *testing.T) {
	for _, testCase := range excludeTestCases {
		assert.Equal(t, testCase.output, Exclude(testCase.source, testCase.exclude), testCase)
	}
}
