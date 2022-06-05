package tokenizer

import (
	"github.com/yanyiwu/gojieba"
)

type Tokenizer struct {
	jieba *gojieba.Jieba
}

func NewTokenizer() *Tokenizer {
	jieba := gojieba.NewJieba()
	return &Tokenizer{
		jieba: jieba,
	}
}

func (j *Tokenizer) Cut(text string) []string {
	if j.jieba == nil {
		j.jieba = gojieba.NewJieba()
	}
	return j.jieba.CutForSearch(text, true)
}

func (j *Tokenizer) Free() {
	j.jieba.Free()
	j = nil
}
