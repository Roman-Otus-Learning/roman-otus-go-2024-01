package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type CacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	mutex    *sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	listItem, isExists := l.items[key]
	if isExists {
		cacheItem := listItem.Value.(*CacheItem)
		cacheItem.value = value
		l.queue.MoveToFront(listItem)

		return isExists
	}

	if l.isCapacityExceeded() {
		l.dropLastItem()
	}

	newCacheItem := &CacheItem{key: key, value: value}
	newListItem := l.queue.PushFront(newCacheItem)
	l.items[key] = newListItem

	return isExists
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	listItem, isExists := l.items[key]
	if !isExists {
		return nil, isExists
	}

	cacheItem := listItem.Value.(*CacheItem)
	l.queue.MoveToFront(listItem)

	return cacheItem.value, isExists
}

func (l *lruCache) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}

func (l *lruCache) dropLastItem() {
	lastListItem := l.queue.Back()
	if lastListItem == nil {
		return
	}

	cacheItem := lastListItem.Value.(*CacheItem)

	delete(l.items, cacheItem.key)
	l.queue.Remove(lastListItem)
}

func (l lruCache) isCapacityExceeded() bool {
	return l.queue.Len() >= l.capacity
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mutex:    &sync.Mutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
