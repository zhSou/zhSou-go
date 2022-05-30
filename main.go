package main

import (
	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/core/dataset"
	menu "github.com/zhSou/zhSou-go/util/menu"
)

var conf = config.Config{
	DataPaths: []string{
		"D:\\data\\wukong_100m_0.dat",
		"D:\\data\\wukong_100m_1.dat",
	},
	DataIndexPaths: []string{
		"D:\\index\\wukong_100m_0.idx",
		"D:\\index\\wukong_100m_1.idx",
	},
	CsvPaths: []string{
		"D:\\input\\wukong_100m_0.csv",
		"D:\\input\\wukong_100m_1.csv",
	},
	InvertedIndexFilePath: "D:\\inverted_index.inv",
	StopWordPath:          "D:\\stop_words.txt",
}

func main() {
	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("csv数据集导入", func() {
		for i := 0; i < 2; i++ {
			dataset.ConvCsvMakeIndexFile(conf.CsvPaths[i], conf.DataPaths[i], conf.DataIndexPaths[i])
		}
	})
	mainMenu.AddItem("构建倒排索引", func() {

	})

	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
