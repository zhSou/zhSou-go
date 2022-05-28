package main

import (
	"encoding/csv"
	"fmt"
	"github.com/flower-corp/rosedb"
	"io"
	"log"
	"os"
)

func u32To4Bytes(n uint32, bs []byte) {
	bs[0] = byte(n >> 24 & 0xff)
	bs[1] = byte(n >> 16 & 0xff)
	bs[2] = byte(n >> 8 & 0xff)
	bs[3] = byte(n & 0xff)
}

/// 将256个csv文件导入至rosedb数据库
/// Failed 实测失败，rosedb会将数据完全预读入内存导致最终内存不足
func main() {
	opts := rosedb.DefaultOptions("D:\\wukong_100m")
	db, err := rosedb.Open(opts)
	defer db.Close()
	if err != nil {
		log.Fatalln("Open rosedb err: ", err)
		return
	}

	var itemId uint32 = 0
	for fileId := 0; fileId < 256; fileId++ {
		inputCsvPath := fmt.Sprintf("D:\\input\\wukong_100m_%d.csv", fileId)
		log.Printf("(已写入%d)开始写入文档：%s", itemId, inputCsvPath)
		file, err := os.Open(inputCsvPath)
		if err != nil {
			log.Println("csv文件打开失败: ", err)
		}
		reader := csv.NewReader(file)
		log.Println("打开成功: ", inputCsvPath)

		// 跳过表头
		_, err = reader.Read()
		if err != nil {
			log.Fatalln("csv文件读取失败：", err)
		}

		for {
			cols, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("csv文件读取失败：", err)
			}
			// 写入数据库，key为itemId，value为[cols[0], cols[1]]
			var keyBs [4]byte
			u32To4Bytes(itemId, keyBs[:])
			db.RPush(keyBs[:], []byte(cols[0]), []byte(cols[1]))
			itemId++
		}

		file.Close()
	}
}
