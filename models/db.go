package models

import (
	"database/sql"

	"github.com/depsea/dep-core/config"
	_ "github.com/lib/pq"
)

// GetDB -
func GetDB() (*sql.DB, error) {
	info, err := config.GetDBInfo()

	if err != nil {
		panic(err)
	}

	return sql.Open("postgres", info)
}
