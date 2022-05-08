package model

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex_NewIndex(t *testing.T) {
	index := NewIndex()
	require.NotNil(t, index)
}

//
//  TestIndex_Add
//  @Description: Test adding a new index
//	测试添加一个新的索引
//  @param t
//
func TestIndex_Add(t *testing.T) {
	index := NewIndex()

	name := "test"
	docId := 1

	// 添加一个不存在的索引
	err := index.Add(name, docId)
	require.NoError(t, err)

	// 分词的索引存在，但是该索引不存在
	err = index.Add(name, docId+1)
	require.NoError(t, err)

	// 添加一个已存在的索引
	err = index.Add(name, docId)
	require.ErrorIs(t, err, ErrIndexExists)

}

//
//  TestIndex_DupAdd
//  @Description: Test adding a duplicate index
//	测试并发添加索引
//  @param t
//
func TestIndex_DupAdd(t *testing.T) {
	index := NewIndex()
	baseName := "test"

	n := 100
	wg := sync.WaitGroup{}
	// 分词不存在添加docId，分词存在添加docId+i
	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			err := index.Add(baseName+strconv.Itoa(i%10), i)
			require.NoError(t, err)
		}()
	}
	wg.Wait()

	// 添加已经存在的docId
	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			err := index.Add(baseName+strconv.Itoa(i%10), i)
			require.ErrorIs(t, err, ErrIndexExists)
		}()
	}
	wg.Wait()
}
