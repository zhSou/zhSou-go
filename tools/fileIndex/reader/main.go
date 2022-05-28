package main

import (
	"encoding/gob"
	"fmt"
	"os"
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

func main() {

	outputGobPath := fmt.Sprintf("D:\\index\\wukong_100m_%d.gob", 0)
	gobFile, _ := os.Open(outputGobPath)
	dec := gob.NewDecoder(gobFile)

	var contentInfo ContentInfo

	dec.Decode(&contentInfo)

	fmt.Println(contentInfo.ItemLength)

	outputPath := fmt.Sprintf("D:\\after\\wukong_100m_%d.dat", 0)
	dataFile, _ := os.Open(outputPath)

	id := 45677
	firstOffset := contentInfo.SeekInfo[id].Offset
	firstUrlSize := contentInfo.SeekInfo[id].UrlLength
	firstTextSize := contentInfo.SeekInfo[id].TextLength
	urlData := make([]byte, firstUrlSize)
	textData := make([]byte, firstTextSize)

	dataFile.ReadAt(urlData, int64(firstOffset))
	dataFile.ReadAt(textData, int64(firstOffset)+int64(firstUrlSize))

	fmt.Println(string(urlData))
	fmt.Println(string(textData))
}
