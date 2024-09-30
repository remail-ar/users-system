package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbpool *pgxpool.Pool

func main() {
	dsn :=
		"host=" + os.Getenv("DB_HOST") +
			" user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" dbname=" + os.Getenv("DB_NAME") +
			" port=" + os.Getenv("DB_PORT") +
			" sslmode=disable TimeZone=UTC"

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Error opening connection to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting sql.DB from gorm.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)                  // Set the maximum number of connections in the idle connection pool
	sqlDB.SetMaxOpenConns(100)                 // Set the maximum number of open connections to the database
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Set the maximum amount of time a connection can remain open

	defer sqlDB.Close()

	// Example usage of a query (replace this with your actual SelectQuery function)
	var result string
	err = db.Raw("SELECT 'Hello, world!'").Scan(&result).Error
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

}
