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
	item := c.items[key]
	queue := &CacheItem{
		Key: key,
		Value: value,
	}
	if item == nil {
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

	item.Value = queue
	c.queue.MoveToFront(item)

	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item := c.items[key]
	if item == nil {
		return nil, false
	}

	c.queue.MoveToFront(item)
	return item.Value.(*CacheItem).Value, true
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
