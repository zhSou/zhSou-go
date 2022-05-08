package model

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valyala/fastrand"
)

//
//  TestDoc_New
//  @Description: Test the New function
//  @param t
//
func TestDoc_New(t *testing.T) {
	doc := NewDoc()
	require.NotNil(t, doc)
}

//
//  TestDoc_Put
//  @Description: Test the Put function
//  @param t
//
func TestDoc_Put(t *testing.T) {
	// 已经被测试过了
	doc := NewDoc()

	id := 1
	name := "foo name"

	doc.Put(id, name)
}

//
//  TestDoc_DupPut
//  @Description: Test the DupPut function
//	测试在并发情况下，是否能够正确的添加数据
//  @param t
//
func TestDoc_DupPut(t *testing.T) {
	// 已经被测试过了
	doc := NewDoc()

	t.Parallel()
	n := 100
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		i := i
		wg.Add(1)
		go func() {
			doc.Put(i, "foo name"+strconv.Itoa(i))
			wg.Done()
		}()
	}
	wg.Wait()
}

//
//  TestDoc_Get
//  @Description: Test the Get function
//  @param t
//
func TestDoc_Get(t *testing.T) {
	doc := NewDoc()
	id := 123
	name := "foo name"
	doc.Put(id, name)

	get, err := doc.Get(id)

	require.NoError(t, err)
	require.Equal(t, name, get)

	get, err = doc.Get(id + 1)
	require.ErrorIs(t, err, ErrDocNotFound)
	require.Equal(t, "", get)

}

//
//  TestDoc_DupGet
//  @Description: Test the DupGet function
//	测试在并发情况下，是否能够正确的获取数据
//  @param t
//
func TestDoc_DupGet(t *testing.T) {
	doc := NewDoc()

	idBegin := 123
	baseName := "foo name"
	n := 100
	wg := sync.WaitGroup{}

	t.Parallel()

	for i := 0; i < n; i++ {
		i := i
		wg.Add(1)
		go func() {
			doc.Put(idBegin+i, baseName+strconv.Itoa(i))
			wg.Done()
		}()
	}
	// 保证之前的Put操作完成
	wg.Wait()

	for i := 0; i < n; i++ {
		i := i
		wg.Add(1)
		go func() {
			get, err := doc.Get(idBegin + i)
			require.NoError(t, err)
			require.Equal(t, baseName+strconv.Itoa(i), get)
			wg.Done()
		}()
	}
	wg.Wait()
}

//
//  TestDoc_RandomGetAndPut
//  @Description: Test the RandomGetAndPut function
//	测试在并发情况下，读写map是否正确
//  @param t
//
func TestDoc_RandomGetAndPut(t *testing.T) {
	doc := NewDoc()
	putN := int(fastrand.Uint32n(100)) + 100
	getN := int(fastrand.Uint32n(100)) + 50
	t.Logf("putN: %d, getN: %d", putN, getN)
	wg := sync.WaitGroup{}

	for i := 0; i < putN; i++ {
		i := i
		wg.Add(1)
		go func() {
			doc.Put(putN+i, strconv.Itoa(i))
			wg.Done()
		}()
	}

	for i := 0; i < getN; i++ {
		i := i
		wg.Add(1)
		go func() {
			_, _ = doc.Get(getN + i)
			wg.Done()
		}()
	}
	wg.Wait()
}
