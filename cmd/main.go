package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo"
)

var pool *pgxpool.Pool

func main() {
	connectionString := "postgres://" +
		os.Getenv("POSTGRES_USER") + ":" +
		os.Getenv("POSTGRES_PASSWORD") + "@" +
		os.Getenv("POSTGRES_ENDPOINT") + ":" +
		os.Getenv("POSTGRES_PORT") + "/" +
		os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	//postgres://postgres:password@localhost:5432
	log.Printf("DEBUG: Connection string is: %s\n", connectionString)
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("v1/login", func(c echo.Context) error {
		username := c.FormValue("user")
		password := c.FormValue("password")
		if username == "" || password == "" {
			return c.String(http.StatusBadRequest, "Missing credentials")
		}
		log.Printf("User: %s Password: %s", username, password)

		var count int
		err := dbpool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&count)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Database error: %v", err))
		}

		if count == 0 {
			// No users exist, insert the new user
			_, err := dbpool.Exec(context.Background(),
				"INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
			if err != nil {
				return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
			}
			return c.String(http.StatusOK, "First user created successfully")
		}

		return c.String(http.StatusOK, "User login successful")
	})
	e.Logger.Fatal(e.Start(":1323"))

	// var greeting string
	// err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)
}
