package main

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type SeekInfo struct {
	Offset     uint32
	UrlLength  uint16
	TextLength uint16
}

type ContentInfo struct {
	ItemLength uint32
	SeekInfo   []SeekInfo
}

func SaveToFile(e any, outputGobPath string) {
	outputFile, _ := os.OpenFile(outputGobPath, os.O_RDWR|os.O_CREATE, 0777)
	defer outputFile.Close()
	enc := gob.NewEncoder(outputFile)
	if err := enc.Encode(e); err != nil {
		log.Fatalln("Gob文件写入失败：", err)
	}
	log.Println("Gob文件成功输出：" + outputGobPath)
}

/// 转化文件，将所有csv文件转化为纯文本文件与索引文件
func ConvertAndMakeFileIndex(inputCsvPath string, outputPath string, outputGobPath string) {
	file, err := os.Open(inputCsvPath)
	if err != nil {
		log.Println("csv文件打开失败: ", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	log.Println("打开成功: ", inputCsvPath)

	outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0777)
	defer outputFile.Close()

	contentInfo := ContentInfo{
		ItemLength: 0,
		SeekInfo:   []SeekInfo{},
	}

	// 字节偏移量
	var offset uint32 = 0
	var outputBytes []byte
	for i := 0; ; i++ {
		cols, err := reader.Read()
		if err == io.EOF {
			fmt.Println(i)
			break
		}
		if err != nil {
			log.Fatalln("csv文件读取失败：", err)
		}
		// 跳过csv表头
		if i == 0 {
			continue
		}

		bs1 := []byte(cols[0])
		bs2 := []byte(cols[1])

		urlLen := uint16(len(bs1))
		textLen := uint16(len(bs2))
		contentInfo.SeekInfo = append(
			contentInfo.SeekInfo,
			SeekInfo{
				offset,
				urlLen,
				textLen,
			})
		outputBytes = append(outputBytes, bs1...)
		outputBytes = append(outputBytes, bs2...)

		offset += uint32(urlLen + textLen)
		contentInfo.ItemLength++
	}
	outputFile.Write(outputBytes)
	outputFile.Sync()
	SaveToFile(contentInfo, outputGobPath)
}

func main() {
	for i := 0; i < 64; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 4; j++ {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				// 输入csv文件
				inputCsvPath := fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", n)
				// 纯文本文件
				outputPath := fmt.Sprintf("D:\\after\\wukong_100m_%d.dat", n)
				// 文件索引文件
				outputGobPath := fmt.Sprintf("D:\\index\\wukong_100m_%d.gob", n)
				ConvertAndMakeFileIndex(inputCsvPath, outputPath, outputGobPath)
				fmt.Println("转化完毕：", inputCsvPath)
			}(i*4 + j)
		}
		wg.Wait()
	}
}
