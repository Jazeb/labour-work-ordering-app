package models;

type Employees struct {
	ID int `gorm:"primary_key"`
	FullName string
	Email string
	Password string
	NationalID int
	SSN int
	Service string
	UserType string
	Role string
	PhoneNumber int
	Area string
	Status bool
}
