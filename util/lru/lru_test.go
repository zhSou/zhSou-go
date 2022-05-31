package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(5)
	assert.NotNil(t, cache)
}

func TestCache_Put(t *testing.T) {
	cache := NewCache(5)
	cache.Put("k1", "v1")
	assert.Equal(t, 1, cache.Len())
	cache.Remove("k1")
	assert.Equal(t, 0, cache.Len())
}

func TestCache_Remove(t *testing.T) {
	cache := NewCache(5)
	assert.Equal(t, 0, cache.Len())

	cache.Remove("k1")
	assert.Equal(t, 0, cache.Len())

	cache.Put("k1", "v1")
	assert.Equal(t, 1, cache.Len())

	cache.Remove("k1")
	assert.Equal(t, 0, cache.Len())
}

func TestCache_Get(t *testing.T) {
	cache := NewCache(1)
	var val Value
	var ok bool
	// 正常情况
	{
		cache.Put("k1", "v1")
		val, ok = cache.Get("k1")
		assert.True(t, ok)
		assert.Equal(t, "v1", val)
		assert.Equal(t, 1, cache.Len())
		cache.Clear()
	}

	// 异常情况
	{
		// 找不到k2
		{
			_, ok = cache.Get("k2")
			assert.False(t, ok)
		}

		// k1被k2置换找不到k1但能找到k2
		{
			cache.Put("k1", "v1")
			cache.Put("k2", "v2")
			_, ok = cache.Get("k1")
			assert.False(t, ok)

			val, ok = cache.Get("k2")
			assert.True(t, ok)
			assert.Equal(t, "v2", val)
			cache.Clear()
		}

	}
}

func TestCache_Len(t *testing.T) {
	cache := NewCache(1)
	assert.Equal(t, 0, cache.Len())

	cache.Put("k1", "v1")
	assert.Equal(t, 1, cache.Len())

	cache.Put("k2", "v2")
	assert.Equal(t, 1, cache.Len())
}

func TestCache_Clear(t *testing.T) {
	cache := NewCache(1)
	assert.Equal(t, 0, cache.Len())

	cache.Put("k1", "v1")
	assert.Equal(t, 1, cache.Len())

	cache.Put("k2", "v2")
	assert.Equal(t, 1, cache.Len())

	cache.Clear()
	assert.Equal(t, 0, cache.Len())
}
