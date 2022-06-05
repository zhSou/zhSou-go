package global

import (
	"github.com/zhSou/zhSou-go/core/tokenizer"
	"sync"
)

var (
	tokenizerOnce sync.Once
	_tokenizer    *tokenizer.Tokenizer
)

func GetTokenizer() *tokenizer.Tokenizer {
	tokenizerOnce.Do(func() {
		_tokenizer = tokenizer.NewTokenizer()
	})
	return _tokenizer
}
