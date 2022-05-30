package main

import (
	"fmt"
	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/core/dataset"
	"github.com/zhSou/zhSou-go/core/invertedindex"
	"github.com/zhSou/zhSou-go/global"
	menu "github.com/zhSou/zhSou-go/util/menu"
	"log"
	"os"
)

func InitConfig(n int) *config.Config {
	var conf = config.Config{
		DataPaths:             []string{},
		DataIndexPaths:        []string{},
		CsvPaths:              []string{},
		InvertedIndexFilePath: "D:\\inverted_index.inv",
		StopWordPath:          "D:\\stop_words.txt",
	}
	for i := 0; i < n; i++ {
		conf.DataPaths = append(conf.DataPaths, fmt.Sprintf("D:\\data\\wukong_100m_%d.dat", i))
		conf.DataIndexPaths = append(conf.DataIndexPaths, fmt.Sprintf("D:\\index\\wukong_100m_%d.idx", i))
		conf.CsvPaths = append(conf.CsvPaths, fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", i))
	}
	return &conf
}

var conf *config.Config
var n = 1

func ImportCsvHandler() {
	// TODO 可以并行优化
	for i := 0; i < n; i++ {
		dataset.ConvCsvMakeIndexFile(conf.CsvPaths[i], conf.DataPaths[i], conf.DataIndexPaths[i])
	}
}

func MakeInvertedIndexHandler() {
	dataReader := global.DataReader
	log.Println("总数据记录数：", dataReader.Len())

	inv := invertedindex.NewInvertedIndex()
	// TODO 可以并行优化
	for i := uint32(0); i < dataReader.Len(); i++ {
		dataRecord, err := dataReader.Read(i)
		if err != nil {
			return
		}
		for _, word := range global.Tokenizer.Cut(dataRecord.Text) {
			if global.StopWordTable.IsStopWord(word) {
				// 停用词过滤
				continue
			}
			inv.Add(word, int(i))
		}
		if i%10000 == 0 {
			log.Printf("当前建立倒排索引百分比 %.2f", float64(i)/float64(dataReader.Len()))
		}
	}
	// 倒排索引升序排序
	inv.Sort()
	invFile, _ := os.OpenFile(conf.InvertedIndexFilePath, os.O_RDWR|os.O_CREATE, 0777)
	defer invFile.Close()
	_ = inv.SaveToDisk(invFile)
	_ = invFile.Sync()
	log.Println("倒排索引文件已存盘：", conf.InvertedIndexFilePath)
}

func QueryInvertedIndexHandler() {
	file, _ := os.Open(conf.InvertedIndexFilePath)
	defer file.Close()
	inv, _ := invertedindex.LoadInvertedIndexFromDisk(file)
	is := inv.Get("手机")
	for i, id := range is {
		if i > 15 {
			break
		}
		record, _ := global.DataReader.Read(uint32(id))
		fmt.Println(record)
	}
}

func main() {
	conf = InitConfig(n)
	global.InitGlobal(conf)

	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("csv数据集导入", ImportCsvHandler)
	mainMenu.AddItem("构建倒排索引", MakeInvertedIndexHandler)
	mainMenu.AddItem("查询倒排索引", QueryInvertedIndexHandler)
	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
