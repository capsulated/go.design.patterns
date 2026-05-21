package storage

import "log"

type Storage struct {
	cache bool
	// redis
}

func NewStorage(cache bool) *Storage {
	return &Storage{cache}
}

func (c *Storage) Store() {
	if c.cache {
		// записываем в кеш
	}

	log.Println("⏳ Store...")
	// ...
}

func (c *Storage) Delete() {
	if c.cache {
		// записываем в кеш
	}

	log.Println("⏳ Delete...")
	// ...
}

// Needed to add cache in storage
//func (*Storage) Cache() {
//	log.Println("⏳ cache data...")
//	// ...
//}
