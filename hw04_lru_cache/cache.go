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

type KeyValue struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	kv := KeyValue{Key: key, Value: value}
	el, exist := c.items[key]

	if exist {
		el.Value = kv
		c.queue.MoveToFront(el)
	} else {
		if c.queue.Len()+1 > c.capacity {
			back := c.queue.Back()
			c.queue.Remove(back)
			delete(c.items, back.Value.(KeyValue).Key)
		}
		c.items[key] = c.queue.PushFront(kv)
	}
	return exist
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	el, exist := c.items[key]
	if exist {
		c.queue.MoveToFront(el)
		return el.Value.(KeyValue).Value, true
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
