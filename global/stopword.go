package global

import (
	"sync"

	"github.com/zhSou/zhSou-go/core/stopword"
)

var (
	stopWordTable     *stopword.Table
	stopWordTableOnce sync.Once
)

func GetStopWordTable() *stopword.Table {
	stopWordTableOnce.Do(func() {
		stopWordTable = stopword.NewTable()
		stopWordTable.LoadStopWord(Config.StopWordPath)
		// 绑定一些额外的停用词规则
		stopWordTable.BindExtraRule(func(word string) bool {
			// 小于等于一字符的为停用词
			if len(word) <= 1 {
				return true
			}
			return false
		})
	})
	return stopWordTable
}
