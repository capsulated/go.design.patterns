package storage

import "log"

type Cacher interface {
	Cache()
}

type CacheStorage struct {
	url            string
	cacheActivated bool
}

func NewCacheStorage(cacheActivated bool) *CacheStorage {
	return &CacheStorage{cacheActivated}
}

func (c *CacheStorage) Store() *data {
	// ...
	if c.cacheActivated {
		// c.url заменяем esacepe
		// !! some data modifications
		c.Cache()
	}
	log.Println("⏳ Saving...")
}

func (c *CacheStorage) StoreEscaped() *data {

}

func (c *CacheStorage) Delete() {
	// ...
	if c.cacheActivated {
		// !! some data modifications
		c.Cache()
	}
}

func (*CacheStorage) Cache() {
	log.Println("⏳ Cache...")
	// ...
}
