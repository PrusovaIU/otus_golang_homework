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
			last_item := c.queue.Back()
			c.queue.Remove(last_item)
			delete(c.items, last_item.Value.(*Item).Key)
		}
		new_item := Item{Key: key, Value: value}
		new_list_item := c.queue.PushFront(&new_item)
		c.items[key] = new_list_item
		flag = false
	}
	return flag
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	var flag bool = false
	var value interface{} = nil
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
