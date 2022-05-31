package global

import (
	"sync"

	"github.com/zhSou/zhSou-go/util/lru"
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
