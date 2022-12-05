package postgres

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/mjur/zippo/pkg/configuration"
	"github.com/pkg/errors"
)

// New create new postgres store.
func New(db *sql.DB) configuration.Store {
	s := &store{
		db: db,
	}

	return s
}

type store struct {
	db *sql.DB
}

func (s *store) GetMainSkus(ctx context.Context, packageName, countryCode string) ([]configuration.MainSku, error) {
	var skus []configuration.MainSku

	rows, err := s.db.QueryContext(ctx, `
	SELECT
		c.package,
		c.country_code,
		c.percentile_min,
		c.percentile_max,
		c.main_sku
	FROM configurations c
	WHERE c.package = $1 AND c.country_code = $2`,
		packageName, countryCode)
	if err != nil {
		e, ok := err.(*pq.Error)
		if ok {
			err = errors.Wrapf(err, "database error %s", e.Code.Class().Name())
		}
		return nil, errors.Wrapf(err, "error querying database")
	}

	defer rows.Close()

	if rows.Err() != nil {
		return nil, errors.Wrapf(err, "rows error")
	}
	for rows.Next() {
		var sku configuration.MainSku
		err := rows.Scan(
			&sku.Package,
			&sku.CountryCode,
			&sku.PercentileMin,
			&sku.PercentileMax,
			&sku.Sku,
		)
		if err != nil {
			return nil, errors.Wrapf(err, "scan error")
		}
		skus = append(skus, sku)
	}
	return skus, nil
}
