package main

import (
	"time"

	"github.com/google/uuid"
)

type Constituent struct {
	Id uuid.UUID
	First_name string
	Last_name string
	Email string
	Phone string
	Address1 string
	Address2 string
	Area string
	City string
	Postcode string
}

type Case struct {
	Id uuid.UUID
	Constituent_id uuid.UUID
	Category string
	Summary string
	Status string
}

type Email struct {
	Id uuid.UUID
	Case_id uuid.UUID
	Datetime time.Time
	From string
	To string
	Cc string
	Subject string
	Content string
	Actioned bool
}

type PageData struct {
	Cases []Case
	Constituents []Constituent
	Emails []Email
}

/*
var emails []Email = []Email{
    {
        Id: uuid.NewRandom(),
        Case_id: uuid.NewRandom(),
        Datetime: "2023-10-10T10:00:00Z",
        From: "john_doe@gmail.com",
        To: "jane_doe@yahoo.com",
        Cc: "",
        Subject: "Test Email",
        Content: "This is a test email",
        Actioned: false,
    },
}
var constituents []Constituent = []Constituent{
	{
		Id: 1,
		First_name: "John",
		Last_name: "Doe",
		Email: "john_doe@gmail.com",
		Phone: "0123456789",
		Address1: "1 Test Street",
		Address2: "",
		Area: "Test Area",
		City: "Test City",
		Postcode: "TE5 7PC",
	},
	{
		Id: 2,
		First_name: "Jane",
		Last_name: "Doe",
		Email: "jane_doe@yahoo.com",
		Phone: "9876543210",
		Address1: "2 Test Street",
		Address2: "",
		Area: "Test Area",
		City: "Test City",
		Postcode: "TE5 7PC",
	},
}

var cases_ []Case = []Case{
	{
		Id: 1,
		Constituent_id: 1,
		Category: "General",
		Summary: "This is a test case",
	},
}

var data = map[string]interface{} {
	"emails": emails,
	"constituents": constituents,
	"cases": cases_,
}
*/