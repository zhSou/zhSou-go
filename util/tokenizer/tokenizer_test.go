package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	tokenizer := NewTokenizer()
	assert.NotNil(t, tokenizer)
}

func TestJiebaTokenizer_Cut(t *testing.T) {
	tokenizer := NewTokenizer()
	cutResult := tokenizer.Cut("执行测试用例")
	assert.Contains(t, cutResult, "执行")
	assert.Contains(t, cutResult, "测试")
	assert.Contains(t, cutResult, "测试用例")
}
