package main

import (
	"log"
)

func main() {
	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", "postgres", "", "127.0.0.1", "5432", "temp")

	// conn, err := pgx.Connect(context.Background(), connStr)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer conn.Close(context.Background())

	// var m1 account
	// row := conn.QueryRow(context.Background(), "SELECT * FROM account LIMIT 1;")
	// err = row.Scan(&m1.Name, &m1.Address, &m1.Birthday)
	// if err != nil {
	// 	fmt.Println("Error Fetching Details:", err)
	// 	return
	// }
	// fmt.Println(m1)
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)

	}
	if err = store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()
}
