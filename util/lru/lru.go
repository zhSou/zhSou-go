package lru

type Key interface{}
type Value interface{}

type Cache struct {
	maxCapacity int
}

// Get / 获取缓存内容
func (c *Cache) Get(key string) (Key, Value) {
	panic("todo")
}

// Put / 放入一个kv
func (c *Cache) Put(key Key, value Value) {
	panic("todo")
}

// Remove / 移除一个指定的key的kv
func (c *Cache) Remove(key Key) {
	panic("todo")
}

// Len / 获取当前已缓存的数量
func (c *Cache) Len() int {
	panic("todo")
}

// Clear / 清空缓存
func (c *Cache) Clear() {
	panic("todo")
}
