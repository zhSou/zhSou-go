package main

import (
	"context"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

const (
	host      = "http://localhost:9200"
	path      = "D:\\noha\\split"
	indexName = "wukong50k_release"
)

var (
	client = &elastic.Client{}
	err    error
)

type Tweet struct {
	Caption string `json:"caption"`
	Url     string `json:"url"`
}

func init() {
	client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
}

func ReadCsv(filepath string) {
	//打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(filepath)
	if err != nil {
		log.Println("csv文件打开失败！")
	}
	defer opencast.Close()

	//创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)
	log.Printf("csv文件打开成功！ %s", filepath)
	for {
		read, err := ReadCsv.Read() //返回切片类型：[chen  hai wei]
		// 结束了
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("csv文件读取失败！,err:", err)
		}
		tweet := Tweet{
			read[0],
			read[1],
		}
		_, err = client.Index().Index(indexName).BodyJson(tweet).Do(context.Background())
		if err != nil {
			log.Fatalln("插入数据失败！")
		}
	}
	log.Printf("导入成功！ %s", filepath)

}

func GetPathFiles(path string) []string {
	paths := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			paths = append(paths, path+"\\"+file.Name())
		}
	}
	return paths
}

func main() {
	paths := GetPathFiles(path)
	for _, path := range paths {
		ReadCsv(path)
	}
}
