package cutpage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testCases = []struct {
		source   []int
		pageId   int
		pageSize int
		res      []int
	}{
		// 按长度为2取第1页，第2页
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   1,
			pageSize: 2,
			res:      []int{5, 2},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   2,
			pageSize: 2,
			res:      []int{8, 3},
		},
		// 按长度为3取第1,2,3页
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   1,
			pageSize: 3,
			res:      []int{5, 2, 8},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   2,
			pageSize: 3,
			res:      []int{3, 1, 2},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   3,
			pageSize: 3,
			res:      []int{},
		},
		// 按长度为4取第1,2,3页
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   1,
			pageSize: 4,
			res:      []int{5, 2, 8, 3},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   2,
			pageSize: 4,
			res:      []int{1, 2},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   3,
			pageSize: 4,
			res:      []int{},
		},
		// 按长度为0取第-1页，第0页，第1页
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   -1,
			pageSize: 0,
			res:      []int{},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   0,
			pageSize: 0,
			res:      []int{},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   1,
			pageSize: 0,
			res:      []int{},
		},
		// 按长度为-1取第-1页，第0页，第1页
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   -1,
			pageSize: -1,
			res:      []int{},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   -1,
			pageSize: 0,
			res:      []int{},
		},
		{
			source:   []int{5, 2, 8, 3, 1, 2},
			pageId:   -1,
			pageSize: 1,
			res:      []int{},
		},
		// 空数据集
		{
			source:   []int{},
			pageId:   1,
			pageSize: 1,
			res:      []int{},
		}, {
			source:   []int{},
			pageId:   22,
			pageSize: 1,
			res:      []int{},
		},
	}
)

func TestCutPage(t *testing.T) {
	for _, testCase := range testCases {
		require.Equal(t, testCase.res, CutPage[int](testCase.source, testCase.pageId, testCase.pageSize), testCase)
	}
}
