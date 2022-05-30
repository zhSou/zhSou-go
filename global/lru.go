package global

import (
	"github.com/zhSou/zhSou-go/util/lru"
	"sync"
)

var (
	searchLru     *lru.Cache
	searchLruOnce sync.Once
)

func GetSearchLru() *lru.Cache {
	searchLruOnce.Do(func() {
		searchLru = lru.NewCache(Config.SearchLruMaxCapacity)
	})
	return searchLru
}
