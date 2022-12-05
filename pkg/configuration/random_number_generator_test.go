package configuration_test

import (
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/mjur/zippo/pkg/configuration"
)

func TestIntn(t *testing.T) {
	nowFunc1 := func() time.Time {
		now, _ := time.Parse(time.RFC1123, time.RFC1123)
		return now
	}
	nowFunc2 := func() time.Time {
		now, _ := time.Parse(time.RFC1123, "Mon, 03 Jan 2006 15:04:05 MST")
		return now
	}
	nowFunc3 := func() time.Time {
		now, _ := time.Parse(time.RFC1123, "Mon, 04 Jan 2006 15:04:05 MST")
		return now
	}

	t.Run("test", func(t *testing.T) {
		is := is.New(t)

		rng := configuration.NewRandomNumberGenerator(nowFunc1)
		res := rng.Intn(100)
		is.Equal(res, 29)

		rng = configuration.NewRandomNumberGenerator(nowFunc2)
		res = rng.Intn(50)
		is.Equal(res, 21)

		rng = configuration.NewRandomNumberGenerator(nowFunc3)
		res = rng.Intn(10)
		is.Equal(res, 1)
	})

}
