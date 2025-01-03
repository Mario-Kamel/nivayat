package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

type makhdoum struct {
	Name     string
	Address  string
	Birthday time.Time
}

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "postgres1234@", "127.0.0.1", "5432", "temp")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var m1 makhdoum
	row := conn.QueryRow(context.Background(), "SELECT * FROM makhdoum LIMIT 1;")
	err = row.Scan(&m1.Name, &m1.Address, &m1.Birthday)
	if err != nil {
		fmt.Println("Error Fetching Details:", err)
		return
	}
	fmt.Println(m1)
}
