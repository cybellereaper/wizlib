package wizlib

import (
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Cache represents a cache for storing fetched data.
type Cache struct {
	data     interface{}
	expiry   time.Time
	duration time.Duration
	mu       sync.RWMutex
}

// NewCache creates a new instance of Cache.
func NewCache(duration time.Duration) *Cache {
	return &Cache{
		duration: duration,
	}
}

// Get retrieves the data from the cache if it is not expired.
func (c *Cache) Get() (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if time.Now().Before(c.expiry) {
		return c.data, true
	}
	return nil, false
}

// Set stores the data in the cache with the specified expiry time.
func (c *Cache) Set(data interface{}, expiry time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = data
	c.expiry = expiry
}

// FetcherCache is a wrapper around DocumentFetcher that adds caching functionality.
type FetcherCache struct {
	DocumentFetcher DocumentFetcher
	Cache           *Cache
}

// Fetch retrieves the HTML document from the cache if available; otherwise, it fetches the document using the wrapped DocumentFetcher and stores it in the cache.
func (c *FetcherCache) Fetch(url string) (*goquery.Document, error) {
	if data, ok := c.Cache.Get(); ok {
		if doc, ok := data.(*goquery.Document); ok {
			return doc, nil
		}
	}

	doc, err := c.DocumentFetcher.Fetch(url)
	if err != nil {
		return nil, err
	}

	c.Cache.Set(doc, time.Now().Add(c.Cache.duration))

	return doc, nil
}
