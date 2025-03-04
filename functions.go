package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func findConstituent(id string) (Constituent, error) {
	constit := Constituent{}
	err := db.QueryRow(`
		SELECT id, first_name, last_name, email, phone, address1, address2, city, postcode 
		FROM constituents
		WHERE id = ?;
	`, id).Scan(&constit.Id, &constit.First_name, &constit.Last_name, &constit.Email, &constit.Phone, &constit.Address1, &constit.Address2, &constit.City, &constit.Postcode)
	if err != nil {
		return Constituent{}, fmt.Errorf("unable to find constituent with id: %s", id)
	}
	return constit, nil
}


func insertConstituent(constituent Constituent, database *sql.DB) error {
	// insert a specific constituent into the 'constituents' database
	_, err := database.Exec(`INSERT INTO constituents (
		first_name, 
		last_name, 
		email, 
		phone, 
		address1, 
		address2, 
		city, 
		postcode
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);`,
		constituent.First_name,
		constituent.Last_name,
		constituent.Email,
		constituent.Phone,
		constituent.Address1,
		constituent.Address2,
		constituent.City,
		constituent.Postcode,
	)
	if err != nil {
		return fmt.Errorf("error: insertConstituent() - %w", err)
	}
	return nil
}


func getConstituents() ([]Constituent, error) {
	// get all constituents from db
	rows, err := db.Query(`
		SELECT id, first_name, last_name, email, phone, address1, address2, city, postcode FROM constituents;
	`)
	if err != nil {
		return []Constituent{}, fmt.Errorf("failed to get constituents from db")
	}
	defer rows.Close()

	var constituents []Constituent
	for rows.Next() {
		var constit Constituent
		err := rows.Scan(&constit.Id, &constit.First_name, &constit.Last_name, &constit.Email, &constit.Phone, &constit.Address1, &constit.Address2, &constit.City, &constit.Postcode)
		if err != nil {
			return []Constituent{}, fmt.Errorf("failed to scan row: %w", err)
		}
		constituents = append(constituents, constit)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return constituents, nil
}


func removeConstituent(id int) error {
	_, err := db.Exec(`
		DELETE FROM constituents WHERE id = ?
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete constituent from db")
	}
	return nil
}


func getAllCases() ([]Case, error) {
	rows, err := db.Query("SELECT id, constituent_id, category, summary FROM cases;")
	if err != nil {
		fmt.Println("Failed to get cases from db")
		return []Case{}, fmt.Errorf("failed to get cases from db")
	}
	cases := []Case{}
	for rows.Next() {
		eachCase := Case{}
		err := rows.Scan(&eachCase.Id, &eachCase.Constituent_id, &eachCase.Category, &eachCase.Summary)
		if err != nil {
			fmt.Println("failed to scan query data into case object")
			return []Case{}, err
		}
		cases = append(cases, eachCase)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return cases, nil
}


func insertCase(c Case) error {
	_, err := db.Exec(`
		INSERT INTO cases (constituent_id, category, summary) VALUES (?, ?, ?)
	`, c.Constituent_id, c.Category, c.Summary)
	if err != nil {
		fmt.Println("failed to insert case into cases table")
		return fmt.Errorf("createCase() failed: %w", err)
	}
	return nil
}


func getConstituentsCases(id string) ([]Case, error) {
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return []Case{}, fmt.Errorf("getConstituentsCases(): failed to convert string to int")
	}
	rows, err := db.Query(`
		SELECT id, constituent_id, category, summary FROM cases WHERE constituent_id = ?; 
	`, int_id)
	if err != nil {
		return []Case{}, fmt.Errorf("getConstituentsCases(): failed to get data from db")
	}
	defer rows.Close()
	cases := []Case{}
	for rows.Next() {
		currentCase := Case{}
		err = rows.Scan(&currentCase.Id, &currentCase.Constituent_id, &currentCase.Category, &currentCase.Summary)
		if err != nil {
			fmt.Println("getConstituentsCases(): error in scanning data into Case object")
			return []Case{}, err
		}
		cases = append(cases, currentCase)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return cases, nil
}

func insertEmail(email Email) error {
	_, err := db.Exec(`
		INSERT INTO emails (email) VALUES (?);
	`, email)
	if err != nil {
		return fmt.Errorf("failed to insert email into db")
	}
	return nil
}