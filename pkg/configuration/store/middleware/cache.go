package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/mjur/zippo/pkg/configuration"
)

type cacheMiddleware struct {
	config *configuration.Config
	next   configuration.Store
	cache  configuration.Cache
}

// NewCacheMiddleware returns a new caching middleware for the configuration store.
func NewCacheMiddleware(config *configuration.Config, next configuration.Store, cache configuration.Cache) configuration.Store {
	m := cacheMiddleware{
		next:   next,
		config: config,
		cache:  cache,
	}

	return &m
}

// GetMainSkus caches the the outputs of the GetMainSku store method.
func (m *cacheMiddleware) GetMainSkus(ctx context.Context, packageName, countryCode string) ([]configuration.MainSku, error) {
	key := fmt.Sprintf("%s:%s", packageName, countryCode)
	cachedSkus, ok := m.cache.Get(key)
	if ok {
		skus, ok := cachedSkus.([]configuration.MainSku)
		if ok {
			return skus, nil
		}
	}

	res, err := m.next.GetMainSkus(ctx, packageName, countryCode)
	m.cache.Set(key, res, time.Second*time.Duration(m.config.TTL))
	return res, err
}
