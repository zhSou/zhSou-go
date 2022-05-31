package global

import (
	"log"
	"os"
	"sync"

	"github.com/zhSou/zhSou-go/core/dict"
)

var (
	dic     *dict.Dict
	dicOnce sync.Once
)

func GetDict() *dict.Dict {
	dicOnce.Do(func() {
		conf := Config
		log.Println("加载字典文件", conf.DictPath)
		dictFile, _ := os.Open(conf.DictPath)
		defer dictFile.Close()
		var err error
		dic, err = dict.Load(dictFile)
		if err != nil {
			log.Fatalln("字典文件加载失败", err)
		}
	})
	return dic
}
