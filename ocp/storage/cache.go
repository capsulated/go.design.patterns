package storage

import "log"

type Cacher interface {
	Cache()
}

type data string

type CacheStorage struct {
	Storage
}

func NewCacheStorage() *CacheStorage {
	return &CacheStorage{}
}

func (c *CacheStorage) Store() (d *data) {
	log.Println("⏳ Saving...")
	return
}

func (c *CacheStorage) Delete() {
	// ...
}

func (*CacheStorage) StoreCache() {
	log.Println("⏳ Cache...")
}

type StorageCacheEvents struct {
	CacheStorage
}
