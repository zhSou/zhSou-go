package global

import (
	"log"
	"os"
	"sync"

	"github.com/zhSou/zhSou-go/core/invertedindex"
)

var (
	invertedIndex     *invertedindex.InvertedIndex
	invertedIndexOnce sync.Once
)

func GetInvertedIndex() *invertedindex.InvertedIndex {
	invertedIndexOnce.Do(func() {
		conf := Config
		log.Println("加载倒排索引文件：", conf.InvertedIndexFilePath)
		invFile, _ := os.Open(conf.InvertedIndexFilePath)
		defer invFile.Close()
		var err error
		invertedIndex, err = invertedindex.LoadInvertedIndexFromDisk(invFile, GetDict())
		if err != nil {
			log.Fatalln("倒排索引文件加载失败", err)
		}
	})
	return invertedIndex
}
