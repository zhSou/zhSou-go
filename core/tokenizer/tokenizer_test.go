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
		// 有效等价类 输入是中文
		{
			input:  "执行测试用例",
			output: []string{"执行", "测试", "试用", "测试用例"},
		},
		// 无效等价类 英文
		{
			input:  "ssdfsdf",
			output: []string{"ssdfsdf"},
		},
		// 特殊符号
		{
			input:  "@#$",
			output: []string{"@", "#", "$"},
		},
		// 数字
		{
			input:  "11212321",
			output: []string{"11212321"},
		},
		// 英文 特殊符号 数字 组合
		{
			input:  "sd@123",
			output: []string{"sd", "@", "123"},
		},
		// 空
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
		assert.Equal(t, testCase.output, cutResult, testCase)
	}
}
