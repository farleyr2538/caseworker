package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"math/rand"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func randomInt(min, max int) int {
	return rand.Intn(max - min) + min
}

// index
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "could not load template", http.StatusInternalServerError)
		fmt.Println("Error parsing template: ", err)
		return
	}
	t.Execute(w, nil)
}

// CONSTITUENTS

func createConstituent(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/layout.html", "templates/create_constituent.html")
	if err != nil {
		http.Error(w, "Could not parse required html file", http.StatusInternalServerError)
		fmt.Println("Error parsing template: ", err)
		return
	}
	t.Execute(w, nil)
}

func viewConstituent(w http.ResponseWriter, r *http.Request) {
	// generate and render a template, inserting the details of a specific constituent

	// get constituent id from request
	r.ParseForm()
	var id string = r.Form.Get("id")
	constituent, err := findConstituent(id)
	// check if function ran correctly
	if err != nil {
		fmt.Println("unable to find constituent with id: ", id)
		http.Error(w, "unable to find constituent with id.", http.StatusInternalServerError)
		return
	}
	constituents := []Constituent{constituent}

	// get constituent's cases
	cases, err := getConstituentsCases(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// put constituent and their cases together into a PageData object
	data := PageData {
		Constituents: constituents,
		Cases: cases,
	}

	// parse HTML files
	t, err := template.ParseFiles("templates/layout.html", "templates/view_constituent.html")
	if err != nil {
		http.Error(w, "Error generating html template (viewConstituent)", http.StatusInternalServerError)
		fmt.Println("Error generating html template (viewConstituent)")
		return
	}
	t.Execute(w, data)
}

func submitConstituent(w http.ResponseWriter, r *http.Request) {

	// get constituent data from request
	parseFormErr := r.ParseForm()
	if parseFormErr != nil {
		http.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}

	// create new constituent instance
	titleCaser := cases.Title(language.English)
	var newConstituent Constituent = Constituent{
		First_name: titleCaser.String(r.Form.Get("first_name")),
		Last_name:  titleCaser.String(r.Form.Get("last_name")),
		Address1:   titleCaser.String(r.Form.Get("address1")),
		Address2:   titleCaser.String(r.Form.Get("address2")),
		City:       titleCaser.String(r.Form.Get("city")),
		Postcode:   cases.Upper(language.English).String(r.Form.Get("postcode")),
		Email:      r.Form.Get("email"),
		Phone:      titleCaser.String(r.Form.Get("phone")),
	}

	// add newConstituent to db
	err := insertConstituent(newConstituent, db)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "failure: insertConstituent()", http.StatusInternalServerError)
		return
	}

	// render template
	t, err := template.ParseFiles("templates/layout.html", "templates/submit.html")
	if err != nil {
		http.Error(w, "error generating html template", http.StatusInternalServerError)
		fmt.Println("error parsing html files")
		return
	}
	t.Execute(w, nil)
}

func viewConstituents(w http.ResponseWriter, r *http.Request) {
	constituents, err := getConstituents()
	if err != nil {
		http.Error(w, "failed to run getConstituents successfully", http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/layout.html", "templates/constituents.html")
	if err != nil {
		http.Error(w, "Could not parse constituents.html", http.StatusInternalServerError)
		fmt.Println("Error parsing template: ", err)
		return
	}

	t.Execute(w, constituents)
}

func deleteConstituent(w http.ResponseWriter, r *http.Request) {

	// get id
	r.ParseForm()
	idString := r.Form.Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("error converting id string to int")
		http.Error(w, "deleteConstituent()", http.StatusInternalServerError)
		return
	}

	// remove constituent with id
	err = removeConstituent(id)
	if err != nil {
		fmt.Println("removeConsittuent() failed")
		http.Error(w, "removeConstituent() failed", http.StatusInternalServerError)
		return
	}

	// get new constituents
	constituents, err := getConstituents()
	if err != nil {
		http.Error(w, "getConstituents() failed", http.StatusInternalServerError)
		return
	}

	// parse new template
	t, err := template.ParseFiles("templates/layout.html", "templates/constituents.html")
	if err != nil {
		http.Error(w, "failed to parse constituents template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, constituents)
}

// CASES

func viewCases(w http.ResponseWriter, r *http.Request) {
	cases, err := getAllCases()
	if err != nil {
		fmt.Println("error while running getAllCases()")
		http.Error(w, "failed to run getAllCases()", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Constituents: nil,
		Cases: cases,
	}

	// build template
	t, err := template.ParseFiles("templates/layout.html", "templates/view_cases.html")
	if err != nil {
		fmt.Println("error creating template")
		http.Error(w, "failed to create template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}


func createCase(w http.ResponseWriter, r *http.Request) {
	// get constituent id from request
	r.ParseForm()
	if r.Form.Get("id") == "" {
		http.Error(w, "id is empty string", http.StatusInternalServerError)
		return
	}
	id := r.Form.Get("id")

	t, err := template.ParseFiles("templates/layout.html", "templates/create_case.html")
	if err != nil {
		http.Error(w, "failed to generate create_case template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, id)
}


func submitCase(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	constituent_id, err := strconv.Atoi(r.Form.Get("constituent_id"))
	if err != nil {
		http.Error(w, "error converting constituent_id to int", http.StatusInternalServerError)
		return
	}
	thisCase := Case{
		Constituent_id: constituent_id,
		Summary: r.Form.Get("summary"),
		Category: r.Form.Get("category"),
	}
	err = insertCase(thisCase)
	if err != nil {
		http.Error(w, "insertCase() failed", http.StatusInternalServerError)
		return
	}

	// take to next page
	t, err := template.ParseFiles("templates/layout.html", "templates/submit_case.html")
	if err != nil {
		http.Error(w, "failed to parse submit_case.html", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}


func viewCase(w http.ResponseWriter, r *http.Request) {
	// get case id
	r.ParseForm()
	id := r.Form.Get("id")
	var constituent_id int
	newCase := Case{}

	// get case details
	err := db.QueryRow(`
		SELECT id, constituent_id, summary, category FROM cases WHERE id = ?
	`, id).Scan(&newCase.Id, &newCase.Constituent_id, &newCase.Summary, &newCase.Category)
	if err != nil {
		http.Error(w, "failed to get cases from db", http.StatusInternalServerError)
		return
	}
	constituent_id = newCase.Constituent_id
	cases := []Case{}
	cases = append(cases, newCase)

	// get constituents details
	c := Constituent{}
	err = db.QueryRow(`
		SELECT id, first_name, last_name, email, phone, address1, address2, city, postcode FROM constituents WHERE id = ?
	`, constituent_id).Scan(&c.Id, &c.First_name, &c.Last_name, &c.Email, &c.Phone, &c.Address1, &c.Address2, &c.City, &c.Postcode)
	if err != nil {
		http.Error(w, "failed to retreive constituent from db", http.StatusInternalServerError)
		return
	}
	constituents := []Constituent{}
	constituents = append(constituents, c)

	// pass into PageData object
	data := PageData{
		Constituents: constituents,
		Cases: cases,
	}

	// render template & pass in PageData object
	t, err := template.ParseFiles("templates/layout.html", "templates/case.html")
	if err != nil {
		http.Error(w, "failed to parse case html file into template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)

}

