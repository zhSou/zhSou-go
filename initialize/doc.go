package initialize

import (
	"github.com/bytedance-basic/zhsou-go/core/model"
	"github.com/bytedance-basic/zhsou-go/global"
)

func InitDoc() error {
	global.Doc = model.NewDoc()

	docD := docData()
	index := 0
	for _, data := range docD {
		global.Doc.Put(index, data)
		index++
	}
	return nil
}

func docData() []string {
	return []string{
		"今年能跑赢96不?备战坦克两项俄军开始选拔参赛队员",
		"13/14赛季 英超第5轮 曼城 vs 曼联 13.09.22",
		"珈黛去黑头系列活动结束",
		"珈黛现在可以换上新的黑头了",
		"济南鲁星灯饰有限公司经过不懈的努力已发展成为现代照明生产行业中重要的一家公司",
	}
}
