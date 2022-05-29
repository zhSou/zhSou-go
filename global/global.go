package global

import (
	"github.com/zhSou/zhSou-go/util/tokenizer"
)

var (
	Tokenizer *tokenizer.Tokenizer
)

func InitGlobal() {
	Tokenizer = tokenizer.NewTokenizer()
}
