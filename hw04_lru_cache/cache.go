package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type Item struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	var flag bool
	if listItem, ok := c.items[key]; ok {
		listItem.Value.(*Item).Value = value
		c.queue.MoveToFront(listItem)
		flag = true
	} else {
		if c.queue.Len() == c.capacity {
			lastItem := c.queue.Back()
			c.queue.Remove(lastItem)
			delete(c.items, lastItem.Value.(*Item).Key)
		}
		newItem := Item{Key: key, Value: value}
		newListItem := c.queue.PushFront(&newItem)
		c.items[key] = newListItem
		flag = false
	}
	return flag
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	flag := false
	var value interface{}
	if listItem, ok := c.items[key]; ok {
		flag = true
		c.queue.MoveToFront(listItem)
		value = listItem.Value.(*Item).Value
	}
	return value, flag
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
