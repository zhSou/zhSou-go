package initialize

import (
	"github.com/yanyiwu/gojieba"

	"github.com/zhSou/zhSou-go/global"
)

func InitJieba() {
	global.Jieba = gojieba.NewJieba()
}
