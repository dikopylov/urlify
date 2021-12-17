package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"sync"
	"urlify/internal/model/infrastructure/config"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func Connect(config *config.Database) *sqlx.DB {
	once.Do(func() {
		source := fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s?sslmode=%s",
			config.Driver, config.User, config.Password, config.Host, config.Port, config.DatabaseName, config.SSLmode,
		)

		db = sqlx.MustConnect(config.Driver, source)
	})

	return db
}

func GetConnection() *sqlx.DB {
	return Connect(nil)
}
