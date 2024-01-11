package database

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

	"github.com/Improwised/GPAT/config"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var dbURL string
var err error

const (
	POSTGRES = "postgres"
	SQLITE3  = "sqlite3"
)

// Connect with database
func Connect(cfg config.DBConfig) (*sql.DB, error) {
	switch cfg.Dialect {
	case POSTGRES:
		return postgresDBConnection(cfg)
	case SQLITE3:
		return sqlite3DBConnection(cfg)
	default:
		return nil, errors.New("no suitable dialect found")
	}
}

func sqlite3DBConnection(cfg config.DBConfig) (*sql.DB, error) {

	if _, err = os.Stat(cfg.SQLiteFilePath); err != nil {
		file, err := os.Create(cfg.SQLiteFilePath)
		if err != nil {
			panic(err)
		}
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}
	db, err = sql.Open(SQLITE3, "./"+cfg.SQLiteFilePath)
	if err != nil {
		return nil, err
	}
	return db, err
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
