package service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/global"
)

func init() {
	rand.Seed(time.Now().Unix())

}

//const pathPrefix = "C:\\Users\\zzq\\Desktop\\新建文件夹 (2)"
const pathPrefix = "D:"

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

func generatorSearchQuery() []string {
	return []string{
		"原神",
		"学校",
		"测试",
		"上海应用技术大学",
		"计算机专业",
		"软件工程专业",
		"搜索引擎开发",
		"上海科技大学",
		"字节跳动测试开发",
		"上海高校",
		"震惊",
		"删库跑路",
		"社区工作人员",
		"接班人",
		"我在画画",
		"六级词汇",
		"软件测试",
		"操作系统教程",
		"中国人",
		"世界杯",
		"葡萄牙夺冠",
		"怎么才能删库跑路",
		"我要删库了",
	}
}

//go test -bench='SearchLight' -benchtime=50s  .
// goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/service
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkSearchLight-12          6876609              8678 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/service       66.660s
func BenchmarkSearchLight(b *testing.B) {
	conf := InitConfigLight()
	global.InitGlobal(conf)
	queries := generatorSearchQuery()
	n := len(queries)
	for i := 0; i < b.N; i++ {
		Search(queries[rand.Intn(n)])
	}
}

// go test -bench='SearchAll' -benchtime=50s  .
// goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/service
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkSearchAll-12              36370           1596453 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/service       112.430s
func BenchmarkSearchAll(b *testing.B) {
	conf := InitConfig()
	global.InitGlobal(conf)
	queries := generatorSearchQuery()
	n := len(queries)
	for i := 0; i < b.N; i++ {
		Search(queries[rand.Intn(n)])
	}
}
