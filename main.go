package main

import (
	"database/sql"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)


// connect to database
var db, dbErr = sql.Open("sqlite3", "./data.db")

func main() {
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()

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

	// add email to database
	http.HandleFunc("/api/add-email", addEmail)

	// log error in case of crash
	log.Fatal(http.ListenAndServe(":8080", nil))
}