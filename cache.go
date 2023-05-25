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

type CacheRaidRepository struct {
	Repository RaidRepository
	Cache      *Cache
}

// GetRaid retrieves a raid by guild ID from the cache if available; otherwise, it fetches the raid using the wrapped RaidRepository and stores it in the cache.
func (c *CacheRaidRepository) GetRaid(guildID string) (*Raid, error) {
	if data, ok := c.Cache.Get(); ok {
		if raid, ok := data.(*Raid); ok {
			return raid, nil
		}
	}

	raid, err := c.Repository.GetRaid(guildID)
	if err != nil {
		return nil, err
	}

	c.Cache.Set(raid, time.Now().Add(c.Cache.duration))

	return raid, nil
}

// SaveRaid saves a raid using the wrapped RaidRepository and updates the cache accordingly.
func (c *CacheRaidRepository) SaveRaid(raid *Raid) error {
	err := c.Repository.SaveRaid(raid)
	if err != nil {
		return err
	}

	c.Cache.Set(raid, time.Now().Add(c.Cache.duration))

	return nil
}

type KioskCache struct {
	Manager *KioskManager
	Cache   *Cache
}

// NewKioskCache creates a new instance of KioskCache.
func NewKioskCache(manager *KioskManager, duration time.Duration) *KioskCache {
	return &KioskCache{
		Manager: manager,
		Cache:   NewCache(duration),
	}
}

// AddItem adds a new item to the kiosk and updates the cache accordingly.
func (kc *KioskCache) AddItem(name, itemType string, item KioskItem) {
	kc.Cache.mu.Lock()
	defer kc.Cache.mu.Unlock()

	kc.Manager.AddItem(name, itemType, item)
	kc.Cache.Set(kc.Manager.GetKiosk(), time.Now().Add(kc.Cache.duration))
}

// GetItem retrieves an item from the kiosk based on its name.
// It first checks the cache and returns the item if found.
// If not found in the cache, it retrieves the item from the underlying KioskManager and updates the cache.
func (kc *KioskCache) GetItem(name string) (KioskItem, error) {
	kc.Cache.mu.RLock()
	defer kc.Cache.mu.RUnlock()

	if data, ok := kc.Cache.Get(); ok {
		if kiosk, ok := data.(Kiosk); ok {
			for _, items := range kiosk.KioskItems {
				if item, ok := items[name]; ok {
					return item, nil
				}
			}
		}
	}

	item, err := kc.Manager.GetItem(name)
	if err != nil {
		return KioskItem{}, err
	}

	kc.Cache.Set(kc.Manager.GetKiosk(), time.Now().Add(kc.Cache.duration))

	return item, nil
}

// RemoveItem removes an item from the kiosk based on its name and item type.
// It updates the cache after removing the item.
func (kc *KioskCache) RemoveItem(name, itemType string) error {
	kc.Cache.mu.Lock()
	defer kc.Cache.mu.Unlock()

	err := kc.Manager.RemoveItem(name, itemType)
	if err != nil {
		return err
	}

	kc.Cache.Set(kc.Manager.GetKiosk(), time.Now().Add(kc.Cache.duration))

	return nil
}

// GetLastUpdated returns the last updated timestamp of the kiosk.
func (kc *KioskCache) GetLastUpdated() int64 {
	kc.Cache.mu.RLock()
	defer kc.Cache.mu.RUnlock()

	if data, ok := kc.Cache.Get(); ok {
		if kiosk, ok := data.(Kiosk); ok {
			return kiosk.LastUpdated
		}
	}

	return 0
}

// GetKiosk returns a copy of the kiosk from the cache.
// If the cache is empty, it retrieves the kiosk from the underlying KioskManager and updates the cache.
func (kc *KioskCache) GetKiosk() (Kiosk, error) {
	kc.Cache.mu.RLock()
	defer kc.Cache.mu.RUnlock()

	if data, ok := kc.Cache.Get(); ok {
		if kiosk, ok := data.(Kiosk); ok {
			return kiosk, nil
		}
	}

	kiosk := kc.Manager.GetKiosk()
	kc.Cache.Set(kiosk, time.Now().Add(kc.Cache.duration))

	return kiosk, nil
}
