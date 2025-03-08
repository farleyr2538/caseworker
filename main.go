package main

import (
	"os"
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func main() {

	// connect to database
	var err error
	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("could not connect to database: ", err)
	}

	// serve dist files
	df := http.FileServer(http.Dir("./dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", df))

	// serve static files
	sf := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", sf))

	// show homescreen
	http.HandleFunc("/", indexHandler)

	// handle other pages
	http.HandleFunc("/constituents", viewConstituents)
	http.HandleFunc("/create-constituent", createConstituent)
	http.HandleFunc("/submit-constituent", submitConstituent)
	http.HandleFunc("/view-constituent", viewConstituent)
	http.HandleFunc("/view-cases", viewCases)
	http.HandleFunc("/delete-constituent", deleteConstituent)
	http.HandleFunc("/create-case", createCase)
	http.HandleFunc("/submit-case", submitCase)
	http.HandleFunc("/case", viewCase)

	// api endpoints
	http.HandleFunc("/api/add-email", addEmail) // add email to database

	// log error in case of crash
	log.Fatal(http.ListenAndServe(":8080", nil))
}