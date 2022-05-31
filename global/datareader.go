package global

import (
	"log"
	"sync"

	"github.com/zhSou/zhSou-go/core/dataset"
)

var (
	dataReader     *dataset.DataReader
	dataReaderOnce sync.Once
)

func GetDataReader() *dataset.DataReader {
	var err error
	dataReaderOnce.Do(func() {
		dataReader, err = dataset.NewDataReader(Config.DataIndexPaths, Config.DataPaths)
		if err != nil {
			log.Fatalln("数据或数据索引加载失败：", err)
		}
	})
	return dataReader
}
