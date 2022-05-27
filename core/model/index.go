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
// Index
//  @Description: 设计的存放倒排索引的数据结构
//  name : [ docId1, docId2, docId3 ]
//
type Index struct {
	data map[string][]int
	sync.RWMutex
}

func NewIndex() *Index {
	return &Index{
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
func (i *Index) GetAll(name string) ([]int, error) {
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
func (i *Index) IsInIndex(name string, docId int) (bool, error) {
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

//
// Add
// 	@Description: 填加一个分词的索引
//	如果该分词不存在，则创建一个新的索引
//	如果该分词存在，但是该文档id不存在，则将该文档id添加到该分词的索引中
//	如果该分词存在，且该文档id存在，则返回 ErrIndexExists
//  @receiver i
//  @param name
//  @param docId
//  @return error
//
func (i *Index) Add(name string, docId int) error {
	find, err := i.IsInIndex(name, docId)

	// 在IsInIndex之后加锁，否则会出现死锁
	i.Lock()
	defer i.Unlock()
	// 如果不存在，则添加
	if errors.Is(err, ErrIndexNotFound) {
		i.data[name] = append(i.data[name], docId)
		return nil
	}
	// 如果之前有索引，并且该文档id已经存在，则不添加
	if find {
		return ErrIndexExists
	}
	i.data[name] = append(i.data[name], docId)
	return err
}
