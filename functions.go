package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	//"github.com/jackc/pgx/v4"
	"github.com/georgysavva/scany/pgxscan"
)

// return a constituent, given a constituent id
func findConstituent(id uuid.UUID) (Constituent, error) {
	constit := Constituent{}
	err := db.QueryRow(context.Background(), `
		SELECT id, first_name, last_name, email, phone, address1, address2, area, city, postcode 
		FROM constituent
		WHERE id = $1;
	`, id).Scan(&constit.Id, &constit.First_name, &constit.Last_name, &constit.Email, &constit.Phone, &constit.Address1, &constit.Address2, &constit.Area, &constit.City, &constit.Postcode)
	if err != nil {
		return Constituent{}, fmt.Errorf("unable to find constituent with id: %s", id)
	}
	return constit, nil
}

// insert a specific constituent into the 'constituents' database
func insertConstituent(constituent Constituent) error {
	commandTag, err := db.Exec(context.Background(), `INSERT INTO constituent (
		first_name, 
		last_name, 
		email, 
		phone, 
		address1, 
		address2, 
		area,
		city, 
		postcode
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
		constituent.First_name,
		constituent.Last_name,
		constituent.Email,
		constituent.Phone,
		constituent.Address1,
		constituent.Address2,
		constituent.Area,
		constituent.City,
		constituent.Postcode,
	)
	if err != nil {
		return fmt.Errorf("error: insertConstituent() - %w", err)
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("error: insertConstituent() - no rows affected")
	}
	return nil
}


func getConstituents() ([]Constituent, error) {
	// get all constituents from db

	var constituents []Constituent
	err := pgxscan.Select(context.Background(), db, &constituents, 
	"SELECT id, first_name, last_name, email, phone, address1, address2, area, city, postcode FROM constituent;")
	if err != nil {
		return []Constituent{}, err
	}
	return constituents, nil

	/*
	rows, err := db.Query(context.Background(), `
		SELECT id, first_name, last_name, email, phone, address1, address2, area, city, postcode FROM constituent;
	`)
	if err != nil {
		return []Constituent{}, fmt.Errorf("failed to get constituents from db")
	}

	constituents, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (Constituent, error) {
		var newConstit Constituent
		err := row.Scan(&newConstit)
		return newConstit, err
	})
	if err != nil {
		return []Constituent{}, err
	}
	return constituents, nil
	*/
}


func removeConstituent(id int) error {
	_, err := db.Exec(context.Background(),`
		DELETE FROM constituent WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete constituent from db")
	}
	return nil
}


func getAllCases() ([]Case, error) {
	cases := []Case{}
	err := pgxscan.Select(context.Background(), db, &cases, "SELECT id, constituent_id, category, summary, status FROM case_;")
	if err != nil {
		return []Case{}, err
	}
	return cases, nil
}


func insertCase(c Case) (case_id uuid.UUID, err error) {
	// return case id of inserted case
	new_id := uuid.New()

	_, err = db.Exec(context.Background(),`
		INSERT INTO case_ (id, constituent_id, category, summary, status) VALUES ($1, $2, $3, $4, $5);
	`, new_id, c.Constituent_id, c.Category, c.Summary, c.Status)
	if err != nil {
		fmt.Println("failed to insert case into cases table")
		return new_id, fmt.Errorf("createCase() failed: %w", err)
	}
	return new_id, nil
}


func getConstituentsCases(constituent_id uuid.UUID) ([]Case, error) {
	cases := []Case{}
	err := pgxscan.Select(context.Background(), db, &cases, "SELECT id, constituent_id, category, summary FROM case_ WHERE constituent_id = $1;", constituent_id)
	if err != nil {
		return []Case{}, err
	}
	return cases, nil
}

func insertEmail(case_id uuid.UUID, datetime time.Time, from string, to string, cc string, subject string, content string, actioned bool) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO email (case_id, datetime, from_, to_, cc, subject, content, actioned) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`, case_id, datetime, from, to, cc, subject, content, actioned)
	if err != nil {
		return fmt.Errorf("failed to insert email into db")
	}
	return nil
}


/*
func referenceGenerator() {
	// generate case reference
}
*/