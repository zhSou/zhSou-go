package global

import (
	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/core/dataset"
	"github.com/zhSou/zhSou-go/core/stopword"
	"github.com/zhSou/zhSou-go/util/tokenizer"
)

var (
	Tokenizer     *tokenizer.Tokenizer
	StopWordTable *stopword.Table
	DataReader    *dataset.DataReader
	Config        *config.Config
)

func InitGlobal(conf *config.Config) {
	Tokenizer = tokenizer.NewTokenizer()
	StopWordTable = stopword.NewTable()
	StopWordTable.LoadStopWord(conf.StopWordPath)
	// 绑定一些额外的停用词规则
	StopWordTable.BindExtraRule(func(word string) bool {
		// 小于等于一字符的为停用词
		if len(word) <= 1 {
			return true
		}
		return false
	})
	DataReader, _ = dataset.NewDataReader(conf.DataIndexPaths, conf.DataPaths)
	Config = conf
}
