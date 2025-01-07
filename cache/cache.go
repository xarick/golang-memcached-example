package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache o'zgaruvchisi
var Cache *memcache.Client

// Memcached ulash
func InitCache(server string) {
	Cache = memcache.New(server)
}

// Keshdan olish
func GetFromCache(key string) (string, error) {
	item, err := Cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

// Keshga yozish
func SetToCache(key, value string, expiration int32) error {
	return Cache.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(value),
		Expiration: expiration,
	})
}
