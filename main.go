package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadENV() {
	// Load .env file (optional error handling)
	if err := godotenv.Load(); err != nil {
		log.Println("[No .env file found] => if you are running this inside container ignore this")
	}
}

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if store.CreateAccount(acc); err != nil {
		log.Fatal(err)
		fmt.Println("something wrong")
		return nil
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "vijay", "mallya", "king999")
}

func main() {
	loadENV()

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seedig the database")
		// seed stuff
		seedAccounts(store)

	}

	server := NewAPIServer(":8001", store)
	server.Run()

}
