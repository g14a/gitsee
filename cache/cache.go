package cache

import (
	"log"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

var RistrettoCache *ristretto.Cache

var once sync.Once

func InitRistrettoCache() *ristretto.Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})

	if err != nil {
		log.Println(err)
	}

	return cache
}

// Cache Set with TTL of 1 day
func Set(key, value interface{}) bool {
	if RistrettoCache.SetWithTTL(key, value, 1, time.Hour*1) {
		return true
	}
	return false
}

func Get(key interface{}) (interface{}, bool) {
	if result, ok := RistrettoCache.Get(key); ok {
		return result, ok
	}
	return nil, false
}

func init() {
	once.Do(func() {
		RistrettoCache = InitRistrettoCache()
	})
}
