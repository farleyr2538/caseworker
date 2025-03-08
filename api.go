package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func addEmail(w http.ResponseWriter, r *http.Request) {

	// get email from request
	r.ParseForm()

	case_id, err := uuid.Parse(r.Form.Get("case_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	datetime := time.Now().UTC()

	// insert email details into db
	err = insertEmail(
		case_id, 
		datetime,
		r.Form.Get("from"),
		r.Form.Get("to"),
		r.Form.Get("cc"),
		r.Form.Get("subject"),
		r.Form.Get("content"),
		false,
	)
	if err != nil {
		http.Error(w, "failed to insert email into db", http.StatusInternalServerError)
		return
	} 

	// return response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
