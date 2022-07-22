// Package glru copied the lru cache from https://github.com/golang/groupcache/blob/master/lru/lru.go and changes to generic.
package glru

import (
	"github.com/ftwp/gpkg/xlist"
)

// Cache is an LRU cache. It is not safe for concurrent access.
type Cache[T any] struct {
	// MaxEntries is the maximum number of cache entries before
	// an item is evicted. Zero means no limit.
	MaxEntries int

	// OnEvicted optionally specifies a callback function to be
	// executed when an entry is purged from the cache.
	OnEvicted func(key Key, value T)

	ll    *xlist.List[*entry[T]]
	cache map[any]*xlist.Element[*entry[T]]
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key any

type entry[T any] struct {
	key   Key
	value T
}

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func New[T any](maxEntries int) *Cache[T] {
	return &Cache[T]{
		MaxEntries: maxEntries,
		ll:         xlist.New[*entry[T]](),
		cache:      make(map[any]*xlist.Element[*entry[T]]),
	}
}

// Add adds a value to the cache.
func (c *Cache[T]) Add(key Key, value T) {
	if c.cache == nil {
		c.cache = make(map[any]*xlist.Element[*entry[T]])
		c.ll = xlist.New[*entry[T]]()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.value = value
		return
	}
	ele := c.ll.PushFront(&entry[T]{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// Get looks up a key's value from the cache.
func (c *Cache[T]) Get(key Key) (value T, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.value, true
	}
	return
}

// Remove removes the provided key from the cache.
func (c *Cache[T]) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache[T]) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache[T]) removeElement(e *xlist.Element[*entry[T]]) {
	c.ll.Remove(e)
	delete(c.cache, e.Value.key)
	if c.OnEvicted != nil {
		c.OnEvicted(e.Value.key, e.Value.value)
	}
}

// Len returns the number of items in the cache.
func (c *Cache[T]) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// Clear purges all stored items from the cache.
func (c *Cache[T]) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			//kv := e.Value.(*entry[T])
			c.OnEvicted(e.Value.key, e.Value.value)
		}
	}
	c.ll = nil
	c.cache = nil
}
