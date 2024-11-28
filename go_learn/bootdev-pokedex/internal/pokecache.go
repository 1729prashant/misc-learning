package pokecache

import (
	"sync"
	"time"
)

// cacheEntry represents an entry in the cache.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache is a structure that holds the cached data and manages concurrency.
type Cache struct {
	entries   map[string]cacheEntry
	mutex     sync.Mutex
	interval  time.Duration
	closeChan chan struct{}
}

// NewCache creates and returns a new Cache instance with a specified interval for reaping stale entries.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:   make(map[string]cacheEntry),
		interval:  interval,
		closeChan: make(chan struct{}),
	}

	// Start the reaping loop in a separate goroutine.
	go cache.reapLoop()

	return cache
}

// Add adds a new entry to the cache.
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves an entry from the cache.
// Returns the value and true if the entry exists; otherwise, nil and false.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

// reapLoop periodically removes stale entries from the cache.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.reapStaleEntries()
		case <-c.closeChan:
			return
		}
	}
}

// reapStaleEntries removes entries that are older than the cache's interval.
func (c *Cache) reapStaleEntries() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.entries, key)
		}
	}
}

// Close stops the reaping loop and cleans up resources.
func (c *Cache) Close() {
	close(c.closeChan)
}
