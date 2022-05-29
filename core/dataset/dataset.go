package dataset

import (
	"encoding/gob"
	"io"
)

// SeekInfo 存放某条记录的寻址信息
type SeekInfo struct {
	Offset     uint32
	UrlLength  uint16
	TextLength uint16
}

// IndexFile 存放索引文件信息
type IndexFile struct {
	ItemLength uint32
	SeekInfo   []SeekInfo
}

// IndexFileSet 数据集索引
type IndexFileSet []IndexFile

// NewIndexFileSet  从文件内容来构造数据集
func NewIndexFileSet(rs []io.Reader) *IndexFileSet {
	var indexFileSet IndexFileSet
	for _, r := range rs {
		var indexFile IndexFile
		gob.NewDecoder(r).Decode(&indexFile)
		indexFileSet = append(indexFileSet, indexFile)
	}
	return &indexFileSet
}

func (*IndexFileSet) Get(id uint32) (fileId int, seekInfo *SeekInfo) {
	// TODO 待实现二分查找算法
	return 0, nil
}
