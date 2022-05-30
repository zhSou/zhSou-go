package stopword

import (
	"bufio"
	"log"
	"os"
)

type Table struct {
	stopWords map[string]struct{}
	extraRule func(word string) bool
}

// LoadStopWord 加载停用词表
func LoadStopWord(filePath string) *Table {
	log.Println("加载停用词表")
	stopWords := make(map[string]struct{})
	// 加载停用词表
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("停用词表加载失败: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		stopWords[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("读取停用词表失败：", err)
	}
	return &Table{
		stopWords: stopWords,
	}
}

// BindExtraRule 绑定额外的停用词规则
func (t *Table) BindExtraRule(extraRule func(word string) bool) {
	t.extraRule = extraRule
}

func (t *Table) IsStopWord(word string) bool {
	if t.extraRule != nil && t.extraRule(word) {
		return true
	}
	_, ok := t.stopWords[word]
	return ok
}
