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
)

func InitConfig() *config.Config {
	var conf = config.Config{
		DataPaths:             []string{},
		DataIndexPaths:        []string{},
		CsvPaths:              []string{},
		InvertedIndexFilePath: "D:\\inverted_index.inv",
		DictPath:              "D:\\dict.dic",
		StopWordPath:          "D:\\stop_words.txt",
		ImportCsvCoroutines:   2,
		PathLength:            4,
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
	dataReader := global.DataReader
	log.Println("总数据记录数：", dataReader.Len())
	dic := dict.NewDict()
	inv := invertedindex.NewInvertedIndex(dic)
	// TODO 可以并行优化
	for i := uint32(0); i < dataReader.Len(); i++ {
		dataRecord, err := dataReader.Read(i)
		if err != nil {
			return
		}
		var words []string
		for _, word := range global.Tokenizer.Cut(dataRecord.Text) {
			if global.StopWordTable.IsStopWord(word) {
				// 停用词过滤
				continue
			}
			words = append(words, word)
		}
		inv.AddWords(words, int(i))

		if i%10000 == 0 {
			log.Printf("当前建立倒排索引百分比 %.2f", float64(i)/float64(dataReader.Len()))
		}
	}
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
	conf := global.Config
	log.Println("加载字典文件")
	dictFile, _ := os.Open(conf.DictPath)
	defer dictFile.Close()
	dic, _ := dict.Load(dictFile)

	log.Println("加载倒排索引文件")
	invFile, _ := os.Open(conf.InvertedIndexFilePath)
	defer invFile.Close()
	inv, _ := invertedindex.LoadInvertedIndexFromDisk(invFile, dic)

	fmt.Printf("请输入查找关键词：")
	var keyword string
	_, _ = fmt.Scanln(&keyword)

	is := inv.Get(keyword)
	for i, id := range is {
		if i >= 10 {
			break
		}
		record, _ := global.DataReader.Read(uint32(id))
		fmt.Println(i, id, record.Text)
	}
}

func ClearHandler() {
	conf := global.Config
	for _, dataPath := range conf.DataPaths {
		filesystem.DeleteFile(dataPath)
	}
	global.DataReader.Close()
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
