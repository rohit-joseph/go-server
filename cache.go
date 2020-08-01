package server

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

var cache *goCache.Cache
var expiry time.Duration

// CreateCache ioitializes the cache
func CreateCache() {
	expiry = 5 * time.Second
	cache = goCache.New(expiry, 10*time.Minute)
}

// PutCache puts pointer to reply message in cache
func PutCache(key []byte, value []byte) {
	k := string(key)
	cache.Set(k, &value, expiry)
}

// GetCache returns pointer to reply message in cache if found, nil if not
func GetCache(key []byte) (*[]byte, bool) {
	k := string(key)
	var ret *[]byte
	if x, found := cache.Get(k); found {
		ret = x.(*[]byte)
		return ret, found
	}
	return nil, false
}
