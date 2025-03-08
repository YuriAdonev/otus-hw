package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type CacheItem struct {
	Key
	Value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	queue := &CacheItem{Key: key, Value: value}
	if item, exists := c.items[key]; exists {
		item.Value = queue
		c.queue.MoveToFront(item)

		return true
	}

	if c.capacity == 0 {
		return false
	}

	if c.capacity == c.queue.Len() {
		lastItem := c.queue.Back()
		if lastItem != nil {
			delete(c.items, lastItem.Value.(*CacheItem).Key)
		}
		c.queue.Remove(lastItem)
	}

	c.items[key] = c.queue.PushFront(queue)

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, exists := c.items[key]; exists {
		c.queue.MoveToFront(item)
		return item.Value.(*CacheItem).Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
