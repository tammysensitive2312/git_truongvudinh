package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	return db, nil
}

func createUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		firstname VARCHAR(50) NOT NULL UNIQUE,
	    lastname VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success create table user")
	return nil
}

func main() {
	db, err := connectDB("root:truong@tcp(localhost:3306)/go_web_example")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	Err := createUserTable(db)
	if Err != nil {
		log.Fatalf("Failed to create user table: %v", Err)
	}

	defer db.Close()
}
