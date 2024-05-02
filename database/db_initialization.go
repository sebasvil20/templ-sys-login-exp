package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitializeDatabase() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=app1 password=root dbname=example_users sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Connected to database\n")
	return db, nil
}
