package service

import (
	"github.com/zhSou/zhSou-go/global"
	"github.com/zhSou/zhSou-go/util/algorithm/set"
)

type searchResult struct {
	Words  []string
	DocIds []int
}

func Search(query string) *searchResult {
	dic := global.GetDict()
	inv := global.GetInvertedIndex()
	tkn := global.GetTokenizer()
	lru := global.GetSearchLru()
	if val, ok := lru.Get(query); ok {
		return val.(*searchResult)
	}

	// 计算查询分词id(过滤所有不存在的分词)
	var queryWordIds []int
	var queryWords []string
	var queryWordDocs [][]int
	for _, kw := range tkn.Cut(query) {
		if id := dic.Get(kw); id != -1 {
			queryWords = append(queryWords, kw)
			queryWordIds = append(queryWordIds, id)
			queryWordDocs = append(queryWordDocs, inv.Get(kw))
		}
	}

	// 计算所有查询分词的交集
	var crossResult []int
	for i := 0; i < len(queryWordDocs); i++ {
		if i == 0 {
			crossResult = queryWordDocs[0]
			continue
		}
		crossResult = set.Cross(crossResult, queryWordDocs[i])
	}

	// 排除所有的过滤词
	ret := &searchResult{
		Words:  queryWords,
		DocIds: crossResult,
	}
	lru.Put(query, ret)
	return ret
}
