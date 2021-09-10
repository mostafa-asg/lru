package lru

import "sync"

// New returns new LRU cache with predefined size
func New(size int) *LRUCache {
	return &LRUCache{
		size: size,
		m:    make(map[string]interface{}),
		ll:   newLinkedList(),
	}
}

// LRUCache represents a LRU cache
type LRUCache struct {
	size  int
	mutex sync.Mutex
	m     map[string]interface{}
	ll    *linklist
}

// Delete removes an item with the given key
func (l *LRUCache) Delete(key string) {
	v, ok := l.m[key]

	if ok {
		l.mutex.Lock()
		l.mutex.Unlock()

		deleteKey := v.(*node).key
		l.ll.delete(v.(*node))
		delete(l.m, deleteKey)
	}
}

// Get retuns an item with the given key
func (l *LRUCache) Get(key string) (interface{}, bool) {
	v, ok := l.m[key]

	if ok {
		return v.(*node).value, ok
	} else {
		return nil, false
	}
}

// Set inserts am item into the cache
func (l *LRUCache) Set(key string, value interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	v, ok := l.m[key]
	if !ok {
		newNode := l.ll.add(key, value)
		l.m[key] = newNode

		if len(l.m) > l.size {
			// evict least recently used
			deleteKey := l.ll.tail.key
			l.ll.delete(l.ll.tail)
			delete(l.m, deleteKey)
		}
	} else {
		lnode := v.(*node)
		lnode.value = value

		if l.ll.head != lnode {
			// add new node to head
			l.ll.delete(lnode)
			l.m[key] = l.ll.add(key, value)
		}
	}
}
