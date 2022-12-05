package cache

import (
	"sync"
	"time"

	"github.com/mjur/zippo/pkg/configuration"
)

func New(timeFunc func() time.Time, data sync.Map) configuration.Cache {
	return &cache{
		timeFunc: timeFunc,
		data:     data,
	}
}

type cache struct {
	timeFunc func() time.Time
	data     sync.Map
}
type element struct {
	value     any
	expiresAt int64
}

func (c *cache) Get(key string) (any, bool) {
	e, ok := c.data.Load(key)
	if !ok {
		return nil, false
	}
	elem, ok := e.(*element)
	if !ok {
		return nil, false
	}
	if elem.expiresAt < c.timeFunc().Unix() {
		c.data.Delete(key)
		return nil, false
	}
	return elem.value, true
}

func (c *cache) Set(key string, val any, expire time.Duration) {
	c.data.Store(key, &element{
		value:     val,
		expiresAt: c.timeFunc().Add(expire).Unix(),
	})
}
