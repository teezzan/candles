// Package database holds database relevant function using SQLX.
package database

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // load MySQL driver
	"github.com/jmoiron/sqlx"
)

// DB is the default database structure holding the SQLX connection.
type DB struct {
	SQL *sqlx.DB
}

// New initializes a new (SQLX) database connection.
func New(
	username string,
	password string,
	host string,
	database string,
) (*DB, error) {
	params := []string{
		"parseTime=true",
		"timeout=30s",      // Note: this is the connection timeout, not the query timeout
		"writeTimeout=30s", // Note: this the timeout for awaiting a block to be written, not the entire query.
		"readTimeout=30s",  // Note: this the timeout for awaiting a block to be read, not the entire result set.
	}

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, host, database, strings.Join(params, "&"))
	db, err := open(uri)
	if err != nil {
		return nil, fmt.Errorf("db initialization failed for main: %w", err)
	}

	return &DB{db}, nil
}

func open(uri string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", uri)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Close implements the io.Closer interface.
func (db *DB) Close() error {
	return db.SQL.Close()
}
