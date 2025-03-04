package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func addEmail(w http.ResponseWriter, r *http.Request) {

	// get email from request
	r.ParseForm()

	// convert case_id
	c_id := r.Form.Get("case_id")
	case_id, err := strconv.Atoi(c_id)
	if err != nil {
		http.Error(w, "failed to convert case_id to int", http.StatusInternalServerError)
		return
	}

	// create new email object
	new_email := Email{
		Id:       randomInt(1, 1000),
		Case_id:  case_id,
		Datetime: r.Form.Get("datetime"),
		From:     r.Form.Get("from"),
		To:       r.Form.Get("to"),
		Cc:       r.Form.Get("cc"),
		Subject:  r.Form.Get("subject"),
		Content:  r.Form.Get("content"),
		Actioned: false,
	}

	// insert email into db

	/* err = insertEmail(email)
	if err != nil {
		http.Error(w, "failed to insert email into db", http.StatusInternalServerError)
		return
	} */

	emails := data["emails"].([]Email)
	emails = append(emails, new_email)
	data["emails"] = emails

	// return response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(new_email)
}
