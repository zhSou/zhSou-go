package global

import (
	"github.com/yanyiwu/gojieba"

	"github.com/bytedance-basic/zhsou-go/core/model"
)

var (
	Index *model.Index
	Doc   *model.Doc
	Jieba *gojieba.Jieba
)
