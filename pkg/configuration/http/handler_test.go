package http_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/matryer/is"
	"github.com/mjur/zippo/pkg/configuration"
	handlers "github.com/mjur/zippo/pkg/configuration/http"
	"github.com/mjur/zippo/pkg/configuration/log"
	"github.com/mjur/zippo/pkg/configuration/mocks"
	"github.com/rs/zerolog"
)

func TestGetMainSku(t *testing.T) {
	cases := []struct {
		it                 string
		packageName        string
		countryCodeHeader  string
		getMainSkuResponse *configuration.MainSku
		getMainSkuError    error

		expectedpackageName string
		expectedCountryCode string
		expectedResult      string
		expectedStatus      int
	}{
		{
			it:                  "calls the generate token service and returns a token",
			packageName:         "mock-packageName",
			expectedpackageName: "mock-packageName",
			countryCodeHeader:   "HR",
			expectedCountryCode: "HR",
			getMainSkuResponse: &configuration.MainSku{
				Package:       "mock-package",
				CountryCode:   "HR",
				PercentileMin: 25,
				PercentileMax: 50,
				Sku:           "mock-sku",
			},
			expectedResult: `"mock-sku"`,
			expectedStatus: http.StatusOK,
		},
		{
			it:                  "returns an error on service error",
			packageName:         "mock-packageName",
			expectedpackageName: "mock-packageName",
			expectedCountryCode: "ZZ",

			getMainSkuError: errors.New("mock-error"),
			expectedResult:  `"Internal server error"`,
			expectedStatus:  http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is := is.New(t)

			service := &mocks.ServiceMock{
				GetMainSkuFunc: func(contextMoqParam context.Context, packageName, countryCode string) (*configuration.MainSku, error) {
					is.Equal(packageName, tc.expectedpackageName)
					is.Equal(countryCode, tc.expectedCountryCode)

					return tc.getMainSkuResponse, tc.getMainSkuError
				},
			}

			req, err := http.NewRequest(http.MethodGet, "/configurations/"+tc.packageName, nil)
			req.Header.Set("X-Appengine-Country", tc.countryCodeHeader)
			req = req.WithContext(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			h := handlers.New(&configuration.Config{
				Log: &log.Logger{
					Log: zerolog.Logger{},
				},
			}, service)

			router := httprouter.New()
			rr := httptest.NewRecorder()

			router.GET("/configurations/:package", h.GetMainSku)
			router.ServeHTTP(rr, req)

			is.Equal(strings.ReplaceAll(rr.Body.String(), "\n", ""), strings.ReplaceAll(tc.expectedResult, "\n", ""))
			is.Equal(rr.Code, tc.expectedStatus)
		})
	}
}
