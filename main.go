package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/zhSou/zhSou-go/api"
	"github.com/zhSou/zhSou-go/core/config"
	"github.com/zhSou/zhSou-go/core/dataset"
	"github.com/zhSou/zhSou-go/core/dict"
	"github.com/zhSou/zhSou-go/core/invertedindex"
	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/service"
	"github.com/zhSou/zhSou-go/util/filesystem"
	"github.com/zhSou/zhSou-go/util/menu"
)

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
	for {
		var keywords string
		fmt.Printf("请输入查找关键词：")
		_, _ = fmt.Scanln(&keywords)

		if keywords == "exit" {
			break
		}

		var searchResult []int
		startTime := time.Now()
		searchResult = service.Search(keywords)
		fmt.Printf("搜索结果：%d条 搜索用时：%.2fs\n", len(searchResult), time.Since(startTime).Seconds())

		fmt.Println("以下预览前十条结果")
		for i, id := range searchResult {
			if i >= 10 {
				break
			}
			record, _ := global.GetDataReader().Read(uint32(id))
			fmt.Println(i, id, record.Text)
		}
		fmt.Printf("全部用时：%.2fs\n", time.Since(startTime).Seconds())
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

func Preload() {
	global.GetDict()
	global.GetStopWordTable()
	global.GetInvertedIndex()
	global.GetTokenizer()
	err := global.GetDataReader().LoadIndexFile()
	if err != nil {
		log.Fatalln("数据索引文件加载失败：", err)
	}
}
func main() {
	var conf *config.Config
	configMenu := menu.NewMenu("请选择配置文件")
	configMenu.AddItem("1亿数据量", func() {
		conf = InitConfig()
		configMenu.StopForNextLoop()
	})
	configMenu.AddItem("30w数据量", func() {
		conf = InitConfigLight()
		configMenu.StopForNextLoop()
	})
	configMenu.Loop()
	global.InitGlobal(conf)

	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("csv数据集导入", ImportCsvHandler)
	mainMenu.AddItem("构建倒排索引", MakeInvertedIndexHandler)
	mainMenu.AddItem("查询", QueryInvertedIndexHandler)
	mainMenu.AddItem("预加载", Preload)
	mainMenu.AddItem("启动查询服务器", api.StartServer)
	mainMenu.AddItem("清理相关文件", ClearHandler)
	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
