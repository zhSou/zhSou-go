package main

import (
	"bufio"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"github.com/yanyiwu/gojieba"
	"io"
	"log"
	"os"
	"sort"
	"sync"
)

var (
	jieba     *gojieba.Jieba
	stopWords map[string]struct{}
)

/// 加载停用词表
func loadStopWord() {
	log.Println("加载停用词表")
	stopWords = make(map[string]struct{})
	// 加载停用词表
	file, err := os.Open("C:\\Users\\i\\Desktop\\zhSou-go\\tools\\index\\stop_words.txt")
	if err != nil {
		log.Fatal("停用词表加载失败: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stopWords[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("读取停用词表失败：", err)
	}
}

/// 判断是否为停用词
func isStopWord(s string) bool {
	_, ok := stopWords[s]
	return ok
}

/// 对csv建立倒排索引
func makeIndex(inputCsvPath string, outputGobPath string) {
	file, err := os.Open(inputCsvPath)
	if err != nil {
		log.Println("csv文件打开失败: ", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	log.Println("打开成功: ", inputCsvPath)

	// 倒排索引
	index := make(map[string][]int)
	for i := 0; ; i++ {
		cols, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("csv文件读取失败：", err)
		}
		doc := cols[1]
		for _, word := range jieba.CutForSearch(doc, true) {
			// 过滤停用词
			if isStopWord(word) {
				continue
			}
			index[word] = append(index[word], i)
		}
	}
	// 索引排序
	for _, ids := range index {
		sort.Ints(ids)
	}

	// 序列化到磁盘
	outputFile, _ := os.OpenFile(outputGobPath, os.O_RDWR|os.O_CREATE, 0777)
	defer outputFile.Close()
	enc := gob.NewEncoder(outputFile)
	if err := enc.Encode(index); err != nil {
		log.Fatalln("Gob文件写入失败：", err)
	}
	log.Println("Gob文件成功输出：" + outputGobPath)
}
func main() {
	// 创建jieba实例
	jieba = gojieba.NewJieba()
	defer jieba.Free()

	// 加载停用词
	loadStopWord()

	// 对一亿条数据分别建立倒排索引并持久化输出为gob文件
	for i := 0; i < 64; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 4; j++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				inputCsvPath := fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", n)
				outputGobPath := fmt.Sprintf("D:\\output\\wukong_100m_%d.gob", n)
				makeIndex(inputCsvPath, outputGobPath)
			}(i*4 + j)
		}
		wg.Wait()
	}
}
