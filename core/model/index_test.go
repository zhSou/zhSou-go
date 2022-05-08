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

//
//  TestIndex_GetAll
//  @Description: Test getting all indexes
//  @param t
//
func TestIndex_GetAll(t *testing.T) {
	index := NewIndex()
	name := "test"
	for i := 0; i < 10; i++ {
		err := index.Add(name, i)
		require.NoError(t, err)
	}
	// 获取有索引的
	docs, err := index.GetAll(name)
	require.NoError(t, err)
	require.Len(t, *docs, 10)

	// 获取没有索引的
	docs, err = index.GetAll(name + "1")
	require.ErrorIs(t, err, ErrIndexNotFound)
	require.Nil(t, docs)
}

//
//  TestIndex_DupGetAll
//  @Description: Test getting all indexes concurrently
//  @param t
//
func TestIndex_DupGetAll(t *testing.T) {
	index := NewIndex()
	baseName := "test"
	n := 100
	wg := sync.WaitGroup{}
	// 给 n/10 个分词添加索引
	for i := 0; i < n/10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for j := 0; j < n; j++ {
				err := index.Add(baseName+strconv.Itoa(i), j)
				require.NoError(t, err)
			}
		}()
	}
	wg.Wait()

	// 已存在的索引
	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			data, err := index.GetAll(baseName + strconv.Itoa(i%10))
			require.NoError(t, err)
			require.Len(t, *data, n)
		}()
	}

	// 不存在的索引
	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			data, err := index.GetAll(baseName + strconv.Itoa(i+10000))
			require.ErrorIs(t, err, ErrIndexNotFound)
			require.Nil(t, data)
		}()
	}
	wg.Wait()
}

//
//  TestIndex_IsInIndex
//  @Description: Test checking if a doc is in an index
//	测试某个文档是否在索引中
//  @param t
//
func TestIndex_IsInIndex(t *testing.T) {
	index := NewIndex()

	name := "test"
	docId := 1

	err := index.Add(name, docId)
	require.NoError(t, err)

	// 有分词 在索引中
	find, err := index.IsInIndex(name, docId)
	require.NoError(t, err)
	require.True(t, find)

	// 有分词，不在索引
	find, err = index.IsInIndex(name, docId+1)
	require.NoError(t, err)
	require.True(t, !find)

	// 没有分词
	find, err = index.IsInIndex(name+"1", docId)
	require.ErrorIs(t, err, ErrIndexNotFound)
	require.True(t, !find)
}

func TestIndex_DuoIsInIndex(t *testing.T) {
	index := NewIndex()
	baseName := "test"
	n := 100
	wg := sync.WaitGroup{}
	// 给 n/10 个分词添加索引
	for i := 0; i < n/10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			for j := 0; j < n; j++ {
				err := index.Add(baseName+strconv.Itoa(i), j)
				require.NoError(t, err)
			}
		}()
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			find, err := index.IsInIndex(baseName+strconv.Itoa(i/10), i)
			require.NoError(t, err)
			require.True(t, find)
		}()
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			find, err := index.IsInIndex(baseName+strconv.Itoa(i/10), i+n)
			require.NoError(t, err)
			require.True(t, !find)
		}()
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			find, err := index.IsInIndex(baseName+strconv.Itoa(i/10+10), i)
			require.ErrorIs(t, err, ErrIndexNotFound)
			require.True(t, !find)
		}()
	}

	wg.Wait()
}
