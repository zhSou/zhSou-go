package binary

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstBigger(t *testing.T) {
	// todo 按照软件测试的要求进行设计测试用例
	s1 := []int{1, 6, 21, 90, 100}
	as := &SliceAccessible[int]{s1}
	// 正常情况
	assert.Equal(t, 0, FindFirstBigger[int](as, -1))
	assert.Equal(t, 0, FindFirstBigger[int](as, 0))

	assert.Equal(t, 1, FindFirstBigger[int](as, 1))
	assert.Equal(t, 1, FindFirstBigger[int](as, 4))
	assert.Equal(t, 3, FindFirstBigger[int](as, 21))

	assert.Equal(t, 4, FindFirstBigger[int](as, 91))
	// 比100,104大的数找不到，返回-1
	assert.Equal(t, -1, FindFirstBigger[int](as, 100))
	assert.Equal(t, -1, FindFirstBigger[int](as, 104))

	// 不存在任何数字，故查找到的索引值恒为-1
	assert.Equal(t, -1, FindFirstBigger[int](NewSliceAccessible([]int{}), 1))
	assert.Equal(t, -1, FindFirstBigger[int](NewSliceAccessible([]int{}), 4))

	s2 := []uint32{1, 3, 6}
	as2 := &SliceAccessible[uint32]{s2}
	assert.Equal(t, 1, FindFirstBigger[uint32](as2, 1))
}

//
//  generateNums
//  @Description: 随机生成 size 大小的 slice
func generateNums(size int) []int {
	// 参数是数组的大小
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums = append(nums, rand.Intn(size))
	}
	sort.Ints(nums)
	return nums
}

func benchmarkFindFirstBigger(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := NewSliceAccessible[int](generateNums(n))
		targetNum := nums.Get(rand.Intn(n))
		_ = FindFirstBigger[int](nums, targetNum)
	}
}

func BenchmarkFindFirstBigger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100) + 10
		nums := NewSliceAccessible[int](generateNums(n))
		targetNum := nums.Get(rand.Intn(n))
		_ = FindFirstBigger[int](nums, targetNum)
	}
}

func BenchmarkFindFirstBigger100(b *testing.B) {
	benchmarkFindFirstBigger(100, b)
}

func BenchmarkFindFirstBigger10000(b *testing.B) {
	benchmarkFindFirstBigger(10000, b)
}

func BenchmarkFindFirstBigger100000(b *testing.B) {
	benchmarkFindFirstBigger(100000, b)
}

func BenchmarkFindFirstBigger1000000(b *testing.B) {
	benchmarkFindFirstBigger(1000000, b)
}

// 执行了 523077次 每次11287ns

// $ go test -bench="FindFirstBigger"  -benchtime=5s .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/algorithm/binary
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkFindFirstBigger-12       523077             11287 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/algorithm/binary 6.071s

// $ go test -bench="FindFirstBigger"  -benchtime=5s -count=3 .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/algorithm/binary
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkFindFirstBigger-12       535316             11557 ns/op
//BenchmarkFindFirstBigger-12       509092             11602 ns/op
//BenchmarkFindFirstBigger-12       528079             11611 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/algorithm/binary 18.623s

// 时间复杂度测试
//$ go test -bench .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/algorithm/binary
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkFindFirstBigger-12               103789             11449 ns/op
//BenchmarkFindFirstBigger100-12             92036             12965 ns/op
//BenchmarkFindFirstBigger10000-12            1022           1181705 ns/op
//BenchmarkFindFirstBigger100000-12             79          14308522 ns/op
//BenchmarkFindFirstBigger1000000-12             7         166170957 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/algorithm/binary 7.481s
