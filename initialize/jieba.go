package initialize

import (
	"github.com/yanyiwu/gojieba"

	"github.com/bytedance-basic/zhsou-go/global"
)

func InitJieba() {
	global.Jieba = gojieba.NewJieba()
}
