package lru

import "container/list"

type Key interface{}
type Value interface{}

type LRU interface {
	Get(key Key) (value Value, ok bool)
	Put(key Key, value Value)
	Remove(key Key)
	Len() int
	Clear()
}

type kv struct {
	key   Key
	value Value
}
type Cache struct {
	maxCapacity int
	linkedList  *list.List
	table       map[Key]*list.Element
}

func NewCache(maxCapacity int) *Cache {
	return &Cache{
		maxCapacity: maxCapacity,
		linkedList:  list.New(),
		table:       make(map[Key]*list.Element),
	}
}

// Get / 获取缓存内容
func (c *Cache) Get(key Key) (value Value, ok bool) {
	if v, ok := c.table[key]; ok {
		c.linkedList.MoveToFront(v)
		return v.Value.(*kv).value, true
	}
	return nil, false
}

// Put / 放入一个kv
func (c *Cache) Put(key Key, value Value) {
	if v, ok := c.table[key]; ok {
		c.linkedList.MoveToFront(v)
		v.Value.(*kv).value = value
	} else {
		c.table[key] = c.linkedList.PushFront(&kv{
			key:   key,
			value: value,
		})
		if c.maxCapacity < c.linkedList.Len() {
			// 淘汰最久没用的
			e := c.linkedList.Back()
			if e != nil {
				c.linkedList.Remove(e)
				kv := e.Value.(*kv)
				delete(c.table, kv.key)
			}
		}
	}
}

// Remove / 移除一个指定的key的kv
func (c *Cache) Remove(key Key) {
	if v, ok := c.table[key]; ok {
		c.linkedList.Remove(v)
		kv := v.Value.(*kv)
		delete(c.table, kv.key)
	}
}

// Len / 获取当前已缓存的数量
func (c *Cache) Len() int {
	return c.linkedList.Len()
}

// Clear / 清空缓存
func (c *Cache) Clear() {
	c.linkedList = list.New()
	c.table = make(map[Key]*list.Element)
}
