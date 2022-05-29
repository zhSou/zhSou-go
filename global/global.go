package global

import (
	"github.com/yanyiwu/gojieba"

	"github.com/zhSou/zhSou-go/core/model"
)

var (
	Index *model.InvertedIndex
	Jieba *gojieba.Jieba
)
