package main

import (
	"log"

	"github.com/coreos/bbolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("ok")
}
