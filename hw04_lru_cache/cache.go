package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item := c.items[key]
	if item == nil {
		if c.capacity == c.queue.Len() {
			lastItem := c.queue.Back()
			for k, i := range c.items {
				if i == lastItem {
					delete(c.items, k)
				}
			}
			c.queue.Remove(lastItem)
		}
		c.items[key] = c.queue.PushFront(value)

		return false
	}

	item.Value = value
	c.queue.MoveToFront(item)

	return true
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item := c.items[key]
	if item == nil {
		return nil, false
	}

	c.queue.MoveToFront(item)
	return item.Value, true
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
