package configuration

import (
	"math/rand"
	"time"
)

//go:generate moq -out ./mocks/rng.go -pkg mocks  . RandomNumberGenerator
type RandomNumberGenerator interface {
	Intn(n int) int
}

func NewRandomNumberGenerator(timeFunc func() time.Time) RandomNumberGenerator {
	return &randomNumberGenerator{
		timeFunc: timeFunc,
	}
}

type randomNumberGenerator struct {
	timeFunc func() time.Time
}

func (r *randomNumberGenerator) Intn(n int) int {
	rand.Seed(r.timeFunc().UnixNano())
	return rand.Intn(n)
}
