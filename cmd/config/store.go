package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/mjur/zippo/pkg/configuration"
	"github.com/mjur/zippo/pkg/configuration/store/postgres"
)

const retryTimeout time.Duration = 5 * time.Second

func NewStore(c configuration.Config) (configuration.Store, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	if host == "" {
		return nil, errors.New("database host is empty")
	}
	if port == "" {
		return nil, errors.New("database port is empty")
	}
	if user == "" {
		return nil, errors.New("database username is empty")
	}
	if password == "" {
		return nil, errors.New("database password is empty")
	}
	if dbname == "" {
		return nil, errors.New("database name is empty")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := retryConn(psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "store connection error")
	}

	err = runMigrations(c, db)
	if err != nil {
		return nil, errors.Wrap(err, "migrate error")
	}
	return postgres.New(db), nil
}

func retryConn(psqlInfo string) (*sql.DB, error) {
	for i := 0; i <= 3; i++ {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			time.Sleep(retryTimeout)
			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}
		time.Sleep(retryTimeout)
	}

	return nil, errors.New("database connection retry exceeded")
}

func runMigrations(c configuration.Config, db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	migrate.SetTable("migrations")
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	c.Log.Infof("Applied %d migrations", n)
	return nil

}
