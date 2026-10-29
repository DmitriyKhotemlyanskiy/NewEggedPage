package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB opens MySQL using the DSN from the DB_DSN environment variable,
// e.g. DB_DSN="user:password@tcp(host:3306)/egged"
func GetDB() (db *sql.DB) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database")
	return db
}
