package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		input  string
		output []string
	}{
		{
			input:  "执行测试用例",
			output: []string{"执行", "测试", "试用", "测试用例"},
		},
		{
			input:  "",
			output: nil,
		},
	}
)

func TestNewTokenizer(t *testing.T) {
	tokenizer := NewTokenizer()
	assert.NotNil(t, tokenizer)
}

func TestJiebaTokenizer_Cut(t *testing.T) {
	tokenizer := NewTokenizer()
	for _, testCase := range testCases {
		cutResult := tokenizer.Cut(testCase.input)
		assert.Equal(t, testCase.output, cutResult)
	}
}
