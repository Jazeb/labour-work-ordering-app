package models;

type User struct {
	ID int `gorm:"primary_key"`
	FirstName string
	LastName string
	Email string
	Password string
	IsAdmin bool
	NationalID int
	SSN int
	Service string
	UserType string
}
