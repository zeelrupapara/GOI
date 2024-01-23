package database

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Improwised/GPAT/config"
)

var db *sql.DB
var dbURL string
var err error

const (
	POSTGRES = "postgres"
)

// Connect with database
func Connect(cfg config.DBConfig) (*sql.DB, error) {
	switch cfg.Dialect {
	case POSTGRES:
		return postgresDBConnection(cfg)
	default:
		return nil, errors.New("no suitable dialect found")
	}
}

func postgresDBConnection(cfg config.DBConfig) (*sql.DB, error) {
	dbURL = "postgres://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + "/" + cfg.Db + "?" + cfg.QueryString
	if db == nil {
		db, err = sql.Open(POSTGRES, dbURL)
		if err != nil {
			return nil, err
		}
		return db, err
	}
	return db, err
}
