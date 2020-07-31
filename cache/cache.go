package cache

import (
	"github.com/dgraph-io/ristretto"
	"log"
	"sync"
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

func init() {
	once.Do(func() {
		RistrettoCache = InitRistrettoCache()
	})
}
