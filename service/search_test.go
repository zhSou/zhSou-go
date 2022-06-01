package service

import (
	"math/rand"
	"testing"
	"time"

	config2 "github.com/zhSou/zhSou-go/config"
	"github.com/zhSou/zhSou-go/global"
)

func init() {
	rand.Seed(time.Now().Unix())

}

func generatorSearchQuery() []string {
	return []string{
		"原神",
		"学校",
		"测试",
		"上海应用技术大学",
		"计算机专业",
		"软件工程专业",
		"搜索引擎开发",
		"上海科技大学",
		"字节跳动测试开发",
		"上海高校",
		"震惊",
		"删库跑路",
		"社区工作人员",
		"接班人",
		"我在画画",
		"六级词汇",
		"软件测试",
		"操作系统教程",
		"中国人",
		"世界杯",
		"葡萄牙夺冠",
		"怎么才能删库跑路",
		"我要删库了",
	}
}

//go test -bench='SearchLight' -benchtime=50s  .
// goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/service
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkSearchLight-12          6876609              8678 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/service       66.660s
func BenchmarkSearchLight(b *testing.B) {
	conf := config2.InitConfigLight()
	global.InitGlobal(conf)
	queries := generatorSearchQuery()
	n := len(queries)
	for i := 0; i < b.N; i++ {
		Search(queries[rand.Intn(n)])
	}
}

// go test -bench='SearchAll' -benchtime=50s  .
// goos: windows
//goarch: amd64
//pkg: github.com/zhSou/zhSou-go/service
//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkSearchAll-12              36370           1596453 ns/op
//PASS
//ok      github.com/zhSou/zhSou-go/service       112.430s
func BenchmarkSearchAll(b *testing.B) {
	conf := config2.InitConfig()
	global.InitGlobal(conf)
	queries := generatorSearchQuery()
	n := len(queries)
	for i := 0; i < b.N; i++ {
		Search(queries[rand.Intn(n)])
	}
}
