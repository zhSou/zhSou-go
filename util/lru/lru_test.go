package lru

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

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

func generatorNewCache(size int, b *testing.B) *Cache {
	cache := NewCache(size)
	require.NotNil(b, cache)
	return cache
}

func benchmarkCachePut(count int, cache *Cache, b *testing.B) {
	for i := 0; i < count; i++ {
		cache.Put("key"+strconv.Itoa(rand.Intn(count)), "val"+strconv.Itoa(rand.Intn(count)))
	}
}

//
//  BenchmarkCachePut100And1
//  @Description: 容量为100 Put 1次
//  @param b
//
func BenchmarkCachePut100And1(b *testing.B) {
	cache := generatorNewCache(100, b)
	for i := 0; i < b.N; i++ {
		benchmarkCachePut(1, cache, b)
	}
}

//
//  BenchmarkCachePut100And100
//  @Description: 容量为100 Put100次
//  @param b
//
func BenchmarkCachePut100And100(b *testing.B) {
	cache := generatorNewCache(100, b)
	for i := 0; i < b.N; i++ {
		benchmarkCachePut(100, cache, b)
	}
}

//
//  BenchmarkCachePut100And10000
//  @Description: 容量为100 put 10000次
//  @param b
//
func BenchmarkCachePut100And10000(b *testing.B) {
	cache := generatorNewCache(100, b)
	for i := 0; i < b.N; i++ {
		benchmarkCachePut(10000, cache, b)
	}
}

//
//  BenchmarkCachePut100And100000
//  @Description: 容量为100 put 100000次
//  @param b
//
func BenchmarkCachePut100And100000(b *testing.B) {
	cache := generatorNewCache(100, b)
	for i := 0; i < b.N; i++ {
		benchmarkCachePut(100000, cache, b)
	}
}

//
//  BenchmarkCachePut100And1000000
//  @Description: 容量为100 put 1000000次
//  @param b
//
func BenchmarkCachePut100And1000000(b *testing.B) {
	cache := generatorNewCache(100, b)
	for i := 0; i < b.N; i++ {
		benchmarkCachePut(1000000, cache, b)
	}
}

// 测试Put的时间复杂度， Put 1次用了842ns
// 						 100 次 里面for循环 100次 时间应该是842*100 但是由于启动可能会有一点时间
//						 10000次 应该是之前的100倍
//$ go test -bench .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/lru
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCachePut100And1-12              1434498               842.0 ns/op
//BenchmarkCachePut100And100-12              32690             37242 ns/op
//BenchmarkCachePut100And10000-12              248           4898855 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/lru      5.633s

// $ go test -bench .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/lru
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCachePut100And1-12              6267732               191.2 ns/op
//BenchmarkCachePut100And100-12              53816             21835 ns/op
//BenchmarkCachePut100And10000-12              246           4811529 ns/op
//BenchmarkCachePut100And100000-12              24          49249317 ns/op
//BenchmarkCachePut100And1000000-12              3         493333600 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/lru      9.851s

func benchmarkCacheGet(count int, cache *Cache, b *testing.B) {
	for i := 0; i < count; i++ {
		cache.Get("key" + strconv.Itoa(rand.Intn(count)))
	}
}

//
//  BenchmarkCacheGet100And1
//  @Description: 容量为100，缓存是满的 测试 Get 1次
//  @param b
//
func BenchmarkCacheGet100And1(b *testing.B) {
	cache := generatorNewCache(100, b)
	benchmarkCachePut(100, cache, b)
	for i := 0; i < b.N; i++ {
		benchmarkCacheGet(1, cache, b)
	}
}

//
//  BenchmarkCacheGet100And100
//  @Description: 容量为100，缓存是满的 测试 Get 1次
//  @param b
//
func BenchmarkCacheGet100And100(b *testing.B) {
	cache := generatorNewCache(100, b)
	benchmarkCachePut(100, cache, b)
	for i := 0; i < b.N; i++ {
		benchmarkCacheGet(100, cache, b)
	}
}

//
//  BenchmarkCacheGet100And10000
//  @Description: 容量为100，缓存是满的 测试 Get 1次
//  @param b
//
func BenchmarkCacheGet100And10000(b *testing.B) {
	cache := generatorNewCache(100, b)
	benchmarkCachePut(100, cache, b)
	for i := 0; i < b.N; i++ {
		benchmarkCacheGet(10000, cache, b)
	}
}

//
//  BenchmarkCacheGet100And100000
//  @Description: 容量为100，缓存是满的 测试 Get 1次
//  @param b
//
func BenchmarkCacheGet100And100000(b *testing.B) {
	cache := generatorNewCache(100, b)
	benchmarkCachePut(100, cache, b)
	for i := 0; i < b.N; i++ {
		benchmarkCacheGet(1000000, cache, b)
	}
}

//
//  BenchmarkCacheGet100And1000000
//  @Description: 容量为100，缓存是满的 测试 Get 1次
//  @param b
//
func BenchmarkCacheGet100And1000000(b *testing.B) {
	cache := generatorNewCache(100, b)
	benchmarkCachePut(100, cache, b)
	for i := 0; i < b.N; i++ {
		benchmarkCacheGet(10000000, cache, b)
	}
}

// $ go test -bench .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/lru
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCachePut100And1-12              6267732               191.2 ns/op
//BenchmarkCachePut100And100-12              53816             21835 ns/op
//BenchmarkCachePut100And10000-12              246           4811529 ns/op
//BenchmarkCachePut100And100000-12              24          49249317 ns/op
//BenchmarkCachePut100And1000000-12              3         493333600 ns/op
//BenchmarkCacheGet100And1-12             21094561                63.11 ns/op
//BenchmarkCacheGet100And100-12             154514              7360 ns/op
//BenchmarkCacheGet100And10000-12             1350            860403 ns/op
//BenchmarkCacheGet100And100000-12              12          97987767 ns/op
//BenchmarkCacheGet100And1000000-12              2         973204150 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/lru      17.017s
