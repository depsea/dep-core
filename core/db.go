package core

import "github.com/go-pg/pg"

// DB -
var DB *pg.DB

// NewDB -
func NewDB() *pg.DB {
	DB = pg.Connect(&pg.Options{
		User:     DBUser,
		Database: DBName,
	})
	return DB
}
