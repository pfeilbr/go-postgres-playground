package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://jhjxzrqwqbqott:qFa7GfRCQCJDHJsDbg4Oa8S1iS@ec2-54-163-249-168.compute-1.amazonaws.com:5432/dat93egs7uul6f")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name FROM person")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Printf("id=%d, name=%s\n", id, name)
	}
}
