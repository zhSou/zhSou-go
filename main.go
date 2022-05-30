package main

import (
	"fmt"
	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/core/dataset"
	"github.com/zhSou/zhSou-go/core/dict"
	"github.com/zhSou/zhSou-go/core/invertedindex"
	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/util/filesystem"
	menu "github.com/zhSou/zhSou-go/util/menu"
	"log"
	"os"
	"sync"
	"time"
)

func InitConfig() *config.Config {
	var conf = config.Config{
		DataPaths:                   []string{},
		DataIndexPaths:              []string{},
		CsvPaths:                    []string{},
		InvertedIndexFilePath:       "D:\\inverted_index.inv",
		DictPath:                    "D:\\dict.dic",
		StopWordPath:                "D:\\stop_words.txt",
		ImportCsvCoroutines:         4,
		MakeInvertedIndexCoroutines: 8,
		PathLength:                  256,
	}
	for i := 0; i < conf.PathLength; i++ {
		conf.DataPaths = append(conf.DataPaths, fmt.Sprintf("D:\\data\\wukong_100m_%d.dat", i))
		conf.DataIndexPaths = append(conf.DataIndexPaths, fmt.Sprintf("D:\\index\\wukong_100m_%d.idx", i))
		conf.CsvPaths = append(conf.CsvPaths, fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", i))
	}
	return &conf
}

func ImportCsvHandler() {
	conf := global.Config
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < conf.ImportCsvCoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range ch {
				dataset.ConvCsvMakeIndexFile(conf.CsvPaths[id], conf.DataPaths[id], conf.DataIndexPaths[id])
			}
		}()
	}
	for i := 0; i < conf.PathLength; i++ {
		ch <- i
	}
	// 关闭通道，可以使得for range退出
	close(ch)
	// 等待开启的所有协程组退出
	wg.Wait()
}

func MakeInvertedIndexHandler() {
	conf := global.Config
	dataReader := global.GetDataReader()
	log.Println("总数据记录数：", dataReader.Len())
	dic := dict.NewDict()                      // 字典
	inv := invertedindex.NewInvertedIndex(dic) // 倒排索引

	ch := make(chan uint32)
	var wg sync.WaitGroup
	for i := 0; i < conf.MakeInvertedIndexCoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range ch {
				dataRecord, err := dataReader.Read(id)
				if err != nil {
					return
				}
				var words []string
				for _, word := range global.GetTokenizer().Cut(dataRecord.Text) {
					if global.GetStopWordTable().IsStopWord(word) {
						// 停用词过滤
						continue
					}
					words = append(words, word)
				}
				inv.AddWords(words, int(id))
			}
		}()
	}

	startTime := time.Now()
	for i := uint32(0); i < dataReader.Len(); i++ {
		ch <- i
		if i%100000 == 0 {
			useTime := time.Since(startTime)
			progress := float64(i) / float64(dataReader.Len())
			remainTime := int((useTime.Seconds() / progress) * (1 - progress))
			log.Printf("任务百分比 %.2f 已用时：%ds预计剩余时间：%ds", progress, int(useTime.Seconds()), remainTime)
		}
	}
	close(ch)
	wg.Wait()
	// 倒排索引升序排序
	inv.Sort()

	// 倒排索引存盘
	invFile, _ := os.OpenFile(conf.InvertedIndexFilePath, os.O_RDWR|os.O_CREATE, 0777)
	defer invFile.Close()
	_ = inv.SaveToDisk(invFile)
	_ = invFile.Sync()
	log.Println("倒排索引文件已存盘：", conf.InvertedIndexFilePath)

	dicFile, _ := os.OpenFile(conf.DictPath, os.O_RDWR|os.O_CREATE, 0777)
	defer dicFile.Close()
	_ = dic.Save(dicFile)
	_ = dicFile.Sync()
	log.Println("词典文件已存盘：", conf.DictPath)
}

func QueryInvertedIndexHandler() {
	inv := global.GetInvertedIndex()
	for {
		var keyword string
		fmt.Printf("请输入查找关键词：")
		_, _ = fmt.Scanln(&keyword)

		if keyword == "exit" {
			break
		}

		is := inv.Get(keyword)

		dataReader := global.GetDataReader()
		for i, id := range is {
			if i >= 10 {
				break
			}
			record, _ := dataReader.Read(uint32(id))
			fmt.Println(i, id, record.Text)
		}
	}
}

func ClearHandler() {
	conf := global.Config
	for _, dataPath := range conf.DataPaths {
		filesystem.DeleteFile(dataPath)
	}
	for _, idxPath := range conf.DataIndexPaths {
		filesystem.DeleteFile(idxPath)
	}
	log.Println("程序即将退出")
	os.Exit(0)
}

func main() {
	global.InitGlobal(InitConfig())
	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("csv数据集导入", ImportCsvHandler)
	mainMenu.AddItem("构建倒排索引", MakeInvertedIndexHandler)
	mainMenu.AddItem("查询倒排索引", QueryInvertedIndexHandler)
	mainMenu.AddItem("启动查询服务器", func() {})
	mainMenu.AddItem("清理相关文件", ClearHandler)
	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
