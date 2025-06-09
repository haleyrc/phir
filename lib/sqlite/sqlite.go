package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Conn struct {
	*sql.DB

	path string
}

func MustOpen(dataSourceName string) *Conn {
	conn, err := Open(dataSourceName)
	if err != nil {
		panic(err)
	}
	return conn
}

func Open(dataSourceName string) (*Conn, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("sqlite: open: %w", err)
	}
	conn := &Conn{
		DB:   db,
		path: dataSourceName,
	}
	return conn, nil
}

func (conn *Conn) Teardown() error {
	if err := conn.Close(); err != nil {
		return fmt.Errorf("conn: teardown: %w", err)
	}
	if err := os.Remove(conn.path); err != nil {
		return fmt.Errorf("conn: teardown: %w", err)
	}
	return nil
}
