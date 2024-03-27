package models

import "time"

type User struct {
	Id                int
	FirstName         string
	LastName          string
	CurrentEmployer   UserEmployer
	PreviousEmployers []UserEmployer
}

type UserEmployer struct {
	Id        int
	UserId    int
	Employer  string
	StartDate time.Time
	EndDate   time.Time
	Provider  *string
}

var Providers = map[string]string{
	"a": "aegon",
	"b": "aviva",
	"c": "legal and general",
	"d": "nest",
	"e": "now pensions",
	"f": "royal london",
	"g": "scottish widows",
	"h": "standard life",
	"i": "the peoples pension",
	"j": "other",
}
