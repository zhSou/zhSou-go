package tokenizer

import (
	"github.com/yanyiwu/gojieba"
)

type Tokenizer interface {
	Cut(text string) []string
}

type JiebaTokenizer struct {
	jieba *gojieba.Jieba
}

func NewTokenizer() Tokenizer {
	jieba := gojieba.NewJieba()
	return &JiebaTokenizer{
		jieba: jieba,
	}
}

func (j *JiebaTokenizer) Cut(text string) []string {
	if j.jieba == nil {
		j.jieba = gojieba.NewJieba()
	}
	return j.jieba.CutForSearch(text, true)
}

func (j *JiebaTokenizer) Free() {
	j.jieba.Free()
	j = nil
}
