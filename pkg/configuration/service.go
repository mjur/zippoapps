package configuration

import (
	"context"

	"github.com/davecgh/go-spew/spew"
)

//go:generate moq -out ./mocks/service.go -pkg mocks  . Service

// Service represents a service.
type Service interface {
	GetMainSku(context.Context, string, string) (*MainSku, error)
}

// New creates a new configuration service.
func New(c *Config, store Store, r RandomNumberGenerator) Service {
	return &service{
		c:     c,
		store: store,
		r:     r,
	}
}

type service struct {
	c     *Config
	store Store
	r     RandomNumberGenerator
}

func (s *service) GetMainSku(ctx context.Context, packageName, countryCode string) (*MainSku, error) {
	skus, err := s.store.GetMainSkus(ctx, packageName, countryCode)
	if err != nil {
		return nil, err
	}

	random := s.r.Intn(100)
	spew.Dump("rand")
	spew.Dump(random)
	var sku *MainSku
	for _, s := range skus {
		if s.PercentileMin <= uint(random) && s.PercentileMax >= uint(random) {
			spew.Dump(s)
			sku = &MainSku{}
			*sku = s
			break
		}
	}

	if sku == nil {
		return nil, NewNotFoundError("no valid sku found")
	}

	return sku, nil
}
