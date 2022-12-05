package configuration

import "context"

//go:generate moq -out ./mocks/store.go -pkg mocks  . Store

// Store defines an interface to store data.
type Store interface {
	GetMainSkus(context.Context, string, string) ([]MainSku, error)
}
