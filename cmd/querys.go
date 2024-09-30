package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SelectQuery(dbpool *pgxpool.Pool) (string, error) {
	var id, email string
	err := dbpool.QueryRow(context.Background(), "select id, primaryemail from USERS").Scan(&id, &email)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(id, email)
	return "False", nil
}
