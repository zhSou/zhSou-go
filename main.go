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

var conf = config.Config{
	DataPaths:             []string{},
	DataIndexPaths:        []string{},
	CsvPaths:              []string{},
	InvertedIndexFilePath: "D:\\inverted_index.inv",
	StopWordPath:          "D:\\stop_words.txt",
}

func InitConfig(n int) {
	for i := 0; i < n; i++ {
		conf.DataPaths = append(conf.DataPaths, fmt.Sprintf("D:\\data\\wukong_100m_%d.dat", i))
		conf.DataIndexPaths = append(conf.DataIndexPaths, fmt.Sprintf("D:\\index\\wukong_100m_%d.idx", i))
		conf.CsvPaths = append(conf.CsvPaths, fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", i))
	}
}

func main() {
	n := 1
	InitConfig(n)
	global.InitGlobal()
	global.StopWordTable.LoadStopWord(conf.StopWordPath)
	// 绑定一些额外的停用词规则
	global.StopWordTable.BindExtraRule(func(word string) bool {
		// 小于等于一字符的为停用词
		if len(word) <= 1 {
			return true
		}
		return false
	})
	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("csv数据集导入", func() {
		// TODO 可以优化成并发执行
		for i := 0; i < n; i++ {
			dataset.ConvCsvMakeIndexFile(conf.CsvPaths[i], conf.DataPaths[i], conf.DataIndexPaths[i])
		}
	})
	mainMenu.AddItem("构建倒排索引", func() {
		dataReader, _ := dataset.NewDataReader(conf.DataIndexPaths, conf.DataPaths)
		log.Println("总数据记录数：", dataReader.Len())

		inv := invertedindex.NewInvertedIndex()
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
	})

	mainMenu.AddItem("查询倒排索引", func() {
		file, _ := os.Open(conf.InvertedIndexFilePath)
		defer file.Close()

	})
	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
