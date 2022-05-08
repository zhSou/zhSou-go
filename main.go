package main

import (
	"github.com/bytedance-basic/zhsou-go/initialize"
)

func main() {
	err := initialize.Init()
	if err != nil {
		panic(err)
	}
}
