package binary

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		source []int // 输入的有序的切片
		target int   // 要找的目标值
		output int   // 期待返回的值
	}{
		// 等价类测试法
		// 有效等价类 ，输入的切片是升序的
		{
			source: []int{9, 10, 20, 30, 40, 50, 60, 70},
			target: 20,
			output: 3,
		},

		// 目标值中没有
		{
			source: []int{9, 10, 20, 30, 40, 50, 60, 70},
			target: 71,
			output: -1,
		},
		// 无效等价类
		// 降序的
		{
			source: []int{70, 60, 50, 40, 30, 20, 10, 9},
			target: 20,
			output: 0,
		},
		// 乱序的
		{
			source: []int{9, 291, 2192, 10, 92912, 993, 99934, 29},
			target: 200,
			output: 4,
		},
		// 空数组
		{
			source: nil,
			target: 10,
			output: -1,
		},

		// 边界值
		{
			source: []int{9, 11, 13, 16, 17, 17, 19, 19, 20, 20},
			target: 8,
			output: 0,
		},
		{
			source: []int{9, 11, 13, 16, 17, 17, 19, 19, 20, 20},
			target: 17,
			output: 6,
		},
		{
			source: []int{9, 11, 13, 16, 17, 17, 19, 19, 20, 20},
			target: 19,
			output: 8,
		},
		{
			source: []int{9, 11, 13, 16, 17, 17, 19, 19, 20, 20},
			target: 20,
			output: -1,
		},

		// 基本路径测试
		// 直接结束
		{
			source: []int{},
			target: 1,
			output: -1,
		},
		// 不会经过循环
		{
			source: []int{1},
			target: 1,
			output: -1,
		},
		{
			source: []int{2},
			target: 1,
			output: 0,
		},
		{
			source: []int{1},
			target: 2,
			output: -1,
		},
		// 经过循环
		{
			source: []int{1, 2, 3},
			target: 1,
			output: 1,
		},
		{
			source: []int{1, 2, 3},
			target: 2,
			output: 2,
		},

		{
			source: []int{1, 2, 3},
			target: 3,
			output: -1,
		},
	}
)

func TestFindFirstBigger(t *testing.T) {
	for _, testCase := range testCases {
		as := &SliceAccessible[int]{testCase.source}
		get := FindFirstBigger[int](as, testCase.target)
		assert.Equal(t, testCase.output, get, testCase)
	}
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
	nums := NewSliceAccessible[int](generateNums(n))
	for i := 0; i < b.N; i++ {
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

//$ go test -bench .
//goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/util/algorithm/binary
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkFindFirstBigger-12                91303             12770 ns/op
//BenchmarkFindFirstBigger100-12          20440143                61.43 ns/op
//BenchmarkFindFirstBigger10000-12         9601911               113.7 ns/op
//BenchmarkFindFirstBigger100000-12        7728529               157.6 ns/op
//BenchmarkFindFirstBigger1000000-12       3631664               281.2 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/util/algorithm/binary 9.662s
