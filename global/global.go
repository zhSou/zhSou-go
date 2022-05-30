package global

import (
	"github.com/zhSou/zhSou-go/core/stopword"
	"github.com/zhSou/zhSou-go/util/tokenizer"
)

var (
	Tokenizer     *tokenizer.Tokenizer
	StopWordTable *stopword.Table
)

func InitGlobal() {
	Tokenizer = tokenizer.NewTokenizer()
	StopWordTable = stopword.NewTable()
}
