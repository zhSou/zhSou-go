package model

import (
	"sync"

	"github.com/pkg/errors"
)

var (
	ErrIndexNotFound = errors.New("index not found")
	ErrIndexExists   = errors.New("index already exists")
)

//
// InvertedIndex
//  @Description: 设计的存放倒排索引的数据结构
//  name : [ docId1, docId2, docId3 ]
//
type InvertedIndex struct {
	data map[string][]int
	sync.RWMutex
}

func NewIndex() *InvertedIndex {
	return &InvertedIndex{
		data: make(map[string][]int),
	}
}

//
// GetAll
//  @Description: 获取所有的索引
//  @receiver i
//  @param name
//  @return []int
//  @return error
//
func (i *InvertedIndex) GetAll(name string) ([]int, error) {
	i.RLock()
	defer i.RUnlock()
	if v, ok := i.data[name]; ok {
		return v, nil
	} else {
		return nil, ErrIndexNotFound
	}
}

//
// IsInIndex
//  @Description: 判断是否在索引中
//	给定一个分词，和一个文档id，判断是否在索引中
//	如果该分词不存在返回 ErrIndexNotFound
//	如果该分词存在，判断是否在该分词的索引中,如果在返回true，否则返回false
//  @receiver i
//  @param name
//  @param docId
//  @return bool
//  @return error
//
func (i *InvertedIndex) IsInIndex(name string, docId int) (bool, error) {
	i.RLock()
	defer i.RUnlock()
	if v, ok := i.data[name]; ok {
		for _, id := range v {
			if id == docId {
				return true, nil
			}
		}
		return false, nil
	} else {
		return false, ErrIndexNotFound
	}
}
