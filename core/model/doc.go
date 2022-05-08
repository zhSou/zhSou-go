package model

import (
	"errors"
	"sync"
)

//
// Doc
//  @Description: 存储文档的结构
//	首先使用map来存，不要求性能，只要方便
//	这是一个加锁的map
//
type Doc struct {
	data map[int]string
	sync.RWMutex
}

var (
	ErrDocNotFound = errors.New("doc not found") // 文档不存在
)

func NewDoc() *Doc {
	return &Doc{
		data: make(map[int]string),
	}
}

//
// Put
//  @Description: 将文档保存至内存
//  @receiver doc
//  @param id
//  @param name
//
func (doc *Doc) Put(id int, name string) {
	doc.Lock()
	defer doc.Unlock()
	doc.data[id] = name
}

//
// Get
//  @Description:
//  @receiver doc
//  @param id
//  @return string
//  @return error
//
func (doc *Doc) Get(id int) (string, error) {
	res, ok := doc.data[id]
	if !ok {
		return "", ErrDocNotFound
	}
	return res, nil
}
