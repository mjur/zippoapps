package configuration_test

import (
	"context"
	"errors"
	"testing"

	"github.com/matryer/is"
	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/mocks"
)

func TestGetMainSku(t *testing.T) {
	// nowFunc := func() time.Time {
	// 	now, _ := time.Parse(time.RFC1123, time.RFC1123)
	// 	return now
	// }
	cases := []struct {
		it                  string
		packageName         string
		countryCode         string
		expectedPackageName string
		expectedCountryCode string
		randomResult        int

		getMainSkusResult []configuration.MainSku
		getMainSkusError  error
		expectedResult    *configuration.MainSku
		expectedError     string
	}{
		{
			it:                  "gets a main sku",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			randomResult:        15,
			getMainSkusResult: []configuration.MainSku{
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
					PercentileMin: 50,
					PercentileMax: 100,
					Sku:           "mock-sku",
				},
			},
			expectedResult: &configuration.MainSku{
				Package:       "mock-package",
				CountryCode:   "mock-code",
				PercentileMin: 0,
				PercentileMax: 25,
				Sku:           "mock-sku",
			},
		},
		{
			it:                  "returns error on no valid sku",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			randomResult:        100,
			getMainSkusResult: []configuration.MainSku{
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
					PercentileMin: 50,
					PercentileMax: 75,
					Sku:           "mock-sku",
				},
			},
			expectedError: "no valid sku found",
		},
		{
			it:                  "returns error on store error",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			randomResult:        100,
			getMainSkusError:    errors.New("mock-error"),
			expectedError:       "mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is := is.New(t)
			store := &mocks.StoreMock{
				GetMainSkusFunc: func(contextMoqParam context.Context, packageName, countryCode string) ([]configuration.MainSku, error) {
					is.Equal(packageName, tc.expectedPackageName)
					is.Equal(countryCode, tc.expectedCountryCode)
					return tc.getMainSkusResult, tc.getMainSkusError
				},
			}
			rng := &mocks.RandomNumberGeneratorMock{
				IntnFunc: func(n int) int {
					return tc.randomResult
				},
			}

			service := configuration.New(&configuration.Config{}, store, rng)
			res, err := service.GetMainSku(context.Background(), tc.packageName, tc.countryCode)
			if err != nil {
				is.Equal(err.Error(), tc.expectedError)
			}
			is.Equal(res, tc.expectedResult)
		})
	}
}
