package middleware_test

import (
	"context"
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/mocks"
	"github.com/mjur/zippo/pkg/configuration/store/middleware"
)

func TestGetMainSku(t *testing.T) {
	cases := []struct {
		it                  string
		packageName         string
		countryCode         string
		expectedPackageName string
		expectedCountryCode string
		getBool             bool
		cacheResult         []configuration.MainSku

		getMainSkuResult []configuration.MainSku
		getMainSkusError error
		expectedResult   []configuration.MainSku
		expectedError    string
	}{
		{
			it:                  "gets the data from cache",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			cacheResult: []configuration.MainSku{
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 0,
					PercentileMax: 25,
					Sku:           "mock-sku",
				},
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 25,
					PercentileMax: 50,
					Sku:           "mock-sku",
				},
			},
			getBool: true,

			expectedResult: []configuration.MainSku{
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 0,
					PercentileMax: 25,
					Sku:           "mock-sku",
				},
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 25,
					PercentileMax: 50,
					Sku:           "mock-sku",
				},
			},
		},
		{
			it:                  "gets the data from db if the cache is empty",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			getMainSkuResult: []configuration.MainSku{
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 0,
					PercentileMax: 25,
					Sku:           "mock-sku",
				},
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 25,
					PercentileMax: 50,
					Sku:           "mock-sku",
				},
			},

			expectedResult: []configuration.MainSku{
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 0,
					PercentileMax: 25,
					Sku:           "mock-sku",
				},
				{
					Package:       "mock-package",
					CountryCode:   "mock-code",
					PercentileMin: 25,
					PercentileMax: 50,
					Sku:           "mock-sku",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is := is.New(t)
			next := &mocks.StoreMock{
				GetMainSkusFunc: func(contextMoqParam context.Context, packageName, countryCode string) ([]configuration.MainSku, error) {
					is.Equal(packageName, tc.expectedPackageName)
					is.Equal(countryCode, tc.expectedCountryCode)
					return tc.getMainSkuResult, tc.getMainSkusError
				},
			}
			cache := &mocks.CacheMock{
				GetFunc: func(key string) (any, bool) {
					return tc.cacheResult, tc.getBool

				},
				SetFunc: func(key string, val any, duration time.Duration) {

				},
			}
			middleware := middleware.NewCacheMiddleware(&configuration.Config{}, next, cache)
			res, err := middleware.GetMainSkus(context.Background(), tc.packageName, tc.countryCode)
			if err != nil {
				is.Equal(err.Error(), tc.expectedError)
			}
			is.Equal(res, tc.expectedResult)
		})
	}
}
