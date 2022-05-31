package dict

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDict(t *testing.T) {
	assert.NotNil(t, NewDict())
}

func TestDict_Put(t *testing.T) {
	dict := NewDict()
	assert.Equal(t, 0, dict.Put("A"))
	assert.Equal(t, 0, dict.Put("A"))
	assert.Equal(t, 1, dict.Put("B"))
	assert.Equal(t, 1, dict.Put("B"))
	assert.Equal(t, 2, dict.Put("C"))
	assert.Equal(t, 0, dict.Put("A"))
}

func TestDict_Get(t *testing.T) {
	dict := NewDict()
	// 正常情况
	assert.Equal(t, 0, dict.Put("A"))
	assert.Equal(t, 0, dict.Get("A"))
	assert.Equal(t, 1, dict.Put("B"))
	assert.Equal(t, 1, dict.Get("B"))
	assert.Equal(t, 0, dict.Get("A"))

	// 无效情况
	assert.Equal(t, -1, dict.Get("C"))
}

func TestDict_Len(t *testing.T) {
	dict := NewDict()
	// 正常情况
	assert.Equal(t, 0, dict.Len())

	assert.Equal(t, 0, dict.Put("A"))
	assert.Equal(t, 1, dict.Len())

	assert.Equal(t, 1, dict.Put("B"))
	assert.Equal(t, 2, dict.Len())

}

func TestDict_Save(t *testing.T) {
	dict := NewDict()
	dict.Put("A")
	dict.Put("B")
	buf := bytes.NewBuffer([]byte{})
	_ = dict.Save(buf)

	loadDict, _ := Load(buf)
	assert.Equal(t, 0, loadDict.Get("A"))
	assert.Equal(t, 1, loadDict.Get("B"))
}
