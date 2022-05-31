package global

import (
	"sync"

	"github.com/zhSou/zhSou-go/util/tokenizer"
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
