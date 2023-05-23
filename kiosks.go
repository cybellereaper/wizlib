package wizlib

import (
	"errors"
	"time"
)

type Kiosk struct {
	Id          string                          `json:"user_id"`
	Title       string                          `json:"title"`
	KioskItems  map[string]map[string]KioskItem `json:"items"`
	LastUpdated int64                           `json:"last_updated"`
}

type KioskItem struct {
	Image       string `json:",omitempty"`
	Stock       int64  `json:"stock"`
	Price       int64  `json:"price"`
	LastUpdated int64  `json:"last_updated"`
}

// Add adds a new item to the kiosk.
func (k *Kiosk) Add(name, itemType string, item KioskItem) {
	k.LastUpdated = time.Now().Unix()

	// Check if the item type already exists in the kiosk.
	if k.KioskItems[itemType] == nil {
		// If not, create a new map for that item type.
		k.KioskItems[itemType] = make(map[string]KioskItem)
	}

	// Check if the item with the given name already exists for the item type.
	if _, ok := k.KioskItems[itemType][name]; !ok {
		// If not, set the stock to a default value of 1 if it's less than or equal to 0.
		if item.Stock <= 0 {
			item.Stock = 1
		}
	}

	// Add or update the item in the kiosk.
	k.KioskItems[itemType][name] = item
}

// Select retrieves an item from the kiosk based on its name.
func (k *Kiosk) Select(name string) (KioskItem, error) {
	// Iterate over each item type in the kiosk.
	for _, items := range k.KioskItems {
		// Check if the item with the given name exists for any item type.
		if item, ok := items[name]; ok {
			// If found, return the item.
			return item, nil
		}
	}

	// If the item is not found, return an error.
	return KioskItem{}, errors.New("item not found")
}

// Remove removes an item from the kiosk based on its name and item type.
func (k *Kiosk) Remove(name, itemType string) error {
	// Check if the item type exists in the kiosk.
	if items, ok := k.KioskItems[itemType]; ok {
		// Check if the item with the given name exists for the item type.
		delete(items, name)

		// If there are no more items for the item type, remove the item type from the kiosk.
		if len(items) == 0 {
			delete(k.KioskItems, itemType)
		}

		return nil
	}

	// If the item is not found, return an error.
	return errors.New("item not found")
}
