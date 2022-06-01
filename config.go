package main

import (
	"fmt"

	"github.com/zhSou/zhSou-go/core/config"
)

//const pathPrefix = "C:\\Users\\zzq\\Desktop\\新建文件夹 (2)"
const pathPrefix = "D:"

func InitConfig() *config.Config {
	var conf = config.Config{
		DataPaths:                   []string{},
		DataIndexPaths:              []string{},
		CsvPaths:                    []string{},
		InvertedIndexFilePath:       pathPrefix + "\\inverted_index.inv",
		DictPath:                    pathPrefix + "\\dict.dic",
		StopWordPath:                pathPrefix + "\\stop_words.txt",
		ImportCsvCoroutines:         4,
		MakeInvertedIndexCoroutines: 8,
		PathLength:                  256,
		SearchLruMaxCapacity:        20,
	}
	for i := 0; i < conf.PathLength; i++ {
		conf.DataPaths = append(conf.DataPaths, fmt.Sprintf("%s\\data\\wukong_100m_%d.dat", pathPrefix, i))
		conf.DataIndexPaths = append(conf.DataIndexPaths, fmt.Sprintf("%s\\index\\wukong_100m_%d.idx", pathPrefix, i))
		conf.CsvPaths = append(conf.CsvPaths, fmt.Sprintf("%s\\input\\wukong_100m_%d.csv", pathPrefix, i))
	}
	return &conf
}

// InitConfigLight 测试时候可以用这个轻量级配置文件，只加载一个数据集
func InitConfigLight() *config.Config {
	var conf = config.Config{
		DataPaths:                   []string{},
		DataIndexPaths:              []string{},
		CsvPaths:                    []string{},
		InvertedIndexFilePath:       pathPrefix + "\\light\\inverted_index.inv",
		DictPath:                    pathPrefix + "\\light\\dict.dic",
		StopWordPath:                pathPrefix + "\\stop_words.txt",
		ImportCsvCoroutines:         4,
		MakeInvertedIndexCoroutines: 8,
		PathLength:                  1,
		SearchLruMaxCapacity:        20,
	}
	for i := 0; i < conf.PathLength; i++ {
		conf.DataPaths = append(conf.DataPaths, fmt.Sprintf("%s\\light\\data\\wukong_100m_%d.dat", pathPrefix, i))
		conf.DataIndexPaths = append(conf.DataIndexPaths, fmt.Sprintf("%s\\light\\index\\wukong_100m_%d.idx", pathPrefix, i))
		conf.CsvPaths = append(conf.CsvPaths, fmt.Sprintf("%s\\input\\wukong_100m_%d.csv", pathPrefix, i))
	}
	return &conf
}
