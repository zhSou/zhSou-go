package menu

import (
	"fmt"
	"strconv"
)

type item struct {
	itemText string
	callback func()
}
type Menu struct {
	title           string
	menuItem        []item
	stopForNextLoop bool
}

func NewMenu(title string) *Menu {
	return &Menu{
		title:           title,
		menuItem:        []item{},
		stopForNextLoop: false,
	}
}

func (m *Menu) AddItem(itemText string, callback func()) {
	m.menuItem = append(m.menuItem, item{itemText, callback})
}

func (m *Menu) AddExitItem(itemText string) {
	m.menuItem = append(m.menuItem, item{itemText, func() {
		m.StopForNextLoop()
	}})
}

func (m *Menu) ShowOnce() {
	fmt.Print("\n\n")
	fmt.Println(m.title)
	for i, item := range m.menuItem {
		fmt.Printf("%d. %s\n", i, item.itemText)
	}
}

func (m *Menu) StopForNextLoop() {
	m.stopForNextLoop = true
}

func (m *Menu) Loop() {
	for !m.stopForNextLoop {
		m.ShowOnce()
		fmt.Print("请选择：")
		var optionStr string
		_, _ = fmt.Scanln(&optionStr)

		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("请输入整数选项！！！")
			continue
		}
		if option < 0 || option >= len(m.menuItem) {
			fmt.Println("请输入合法的选项编号！！！")
			continue
		}
		m.menuItem[option].callback()
	}
}
