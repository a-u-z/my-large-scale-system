package pg

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	host = "postgres" // 建立 images 的時候用這個
	// host     = "localhost" // 在本機跑的時候用這個
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

type Database struct {
	db *sql.DB
}

func NewDB() (*Database, error) {
	var db *sql.DB
	var err error
	maxRetries := 10

	for i := 1; i <= maxRetries; i++ {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		log.Printf("Attempt %d: here is connStr:%+v", i, connStr)

		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Error connecting to the database: %v", err)
			if i < maxRetries {
				log.Printf("Retrying in 5 seconds...")
				time.Sleep(5 * time.Second)
			}
		} else {
			break
		}
	}

	if err != nil {
		log.Fatal("Exceeded maximum retries, unable to connect to the database.")
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}
