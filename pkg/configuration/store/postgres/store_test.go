package postgres_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/matryer/is"
	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/store/postgres"
)

func TestGetMainSkus(t *testing.T) {
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
		sqlError            error
		expectedResult      []configuration.MainSku
		expectedError       string
	}{
		{
			it:                  "gets a main sku",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
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
			it:                  "returns error on sql error",
			packageName:         "mock-package",
			countryCode:         "mock-code",
			expectedPackageName: "mock-package",
			expectedCountryCode: "mock-code",
			sqlError:            errors.New("mock-error"),
			expectedError:       "error querying database: mock-error",
		},
	}

	for _, tc := range cases {
		t.Run(tc.it, func(t *testing.T) {
			is := is.New(t)
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()
			query := mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT
					c.package,
					c.country_code,
					c.percentile_min,
					c.percentile_max,
					c.main_sku
				FROM configurations c
				WHERE c.package = $1 AND c.country_code = $2`)).
				WithArgs(tc.expectedPackageName, tc.expectedCountryCode)
			mockRows := sqlmock.NewRows([]string{
				"package", "country_code", "percentile_min", "percentile_max", "main_sku",
			}).AddRow("mock-package", "mock-code", 0, 25, "mock-sku").
				AddRow("mock-package", "mock-code", 25, 50, "mock-sku")

			query.WillReturnRows(mockRows)
			query.WillReturnError(tc.sqlError)

			store := postgres.New(db)
			res, err := store.GetMainSkus(context.Background(), tc.packageName, tc.countryCode)
			if err != nil {
				is.Equal(err.Error(), tc.expectedError)
			}
			is.Equal(res, tc.expectedResult)
		})
	}
}
