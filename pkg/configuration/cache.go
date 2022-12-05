package configuration

import "time"

//go:generate moq -out ./mocks/cache.go -pkg mocks  . Cache
type Cache interface {
	Get(string) (any, bool)
	Set(string, any, time.Duration)
}
