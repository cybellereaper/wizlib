package wizlib

import (
	"errors"
	"time"
)

type Kiosk struct {
	Id          string                          `bson:"user_id"`
	Title       string                          `bson:"title"`
	KioskItems  map[string]map[string]KioskItem `bson:"items"`
	LastUpdated int64                           `bson:"last_updated"`
}

type KioskItem struct {
	Image       string `bson:",omitempty"`
	Stock       int64  `bson:"stock"`
	Price       int64  `bson:"price"`
	LastUpdated int64  `bson:"last_updated"`
}

func (k *Kiosk) Add(name, itemType string, item KioskItem) {
	k.LastUpdated = time.Now().Unix()
	if k.KioskItems[itemType] == nil {
		k.KioskItems[itemType] = make(map[string]KioskItem)
	}
	if _, ok := k.KioskItems[itemType][name]; !ok {
		if item.Stock <= 0 {
			item.Stock = 1
		}
	}
	k.KioskItems[itemType][name] = item
}

func (k *Kiosk) Select(name string) (KioskItem, error) {
	for _, items := range k.KioskItems {
		if item, ok := items[name]; ok {
			return item, nil
		}
	}
	return KioskItem{}, errors.New("item not found")
}

func (k *Kiosk) Remove(name, itemType string) error {
	if items, ok := k.KioskItems[itemType]; ok {
		delete(items, name)
		if len(items) == 0 {
			delete(k.KioskItems, itemType)
		}
		return nil
	}
	return errors.New("item not found")
}
