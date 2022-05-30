package invertedindex

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/zhSou/zhSou-go/core/dict"
	"testing"
)

func TestNewInvertedIndex(t *testing.T) {
	assert.NotNil(t, NewInvertedIndex(dict.NewDict()))
}

func TestInvertedIndex_Add(t *testing.T) {
	ii := NewInvertedIndex(dict.NewDict())
	ii.Add("w1", 5)
	ii.Add("w1", 8)
	ii.Add("w1", 1)
	ii.Add("w1", 6)
	assert.Equal(t, []int{5, 8, 1, 6}, ii.Get("w1"))
}

func TestInvertedIndex_Sort(t *testing.T) {
	ii := NewInvertedIndex(dict.NewDict())
	ii.Add("w1", 5)
	ii.Add("w1", 8)
	ii.Add("w1", 1)
	ii.Add("w1", 6)
	ii.Sort()
	assert.Equal(t, []int{1, 5, 6, 8}, ii.Get("w1"))
}

func TestInvertedIndex_SaveToDisk(t *testing.T) {
	ii := NewInvertedIndex(dict.NewDict())
	ii.Add("w1", 5)
	ii.Add("w1", 8)
	ii.Add("w1", 1)
	ii.Add("w1", 6)
	ii.Sort()

	buf := bytes.NewBuffer([]byte{})
	err := ii.SaveToDisk(buf)
	if err != nil {
		panic(err)
	}
	ifd, err := LoadInvertedIndexFromDisk(buf)
	if err != nil {
		return
	}
	assert.Equal(t, ii.Data, ifd.Data)
}
