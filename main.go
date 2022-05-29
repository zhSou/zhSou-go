package main

import (
	menu "github.com/zhSou/zhSou-go/util/menu"
)

func main() {
	mainMenu := menu.NewMenu("主菜单")
	mainMenu.AddItem("构建倒排索引", func() {

	})

	mainMenu.AddExitItem("退出")
	mainMenu.Loop()
}
