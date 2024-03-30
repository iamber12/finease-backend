package models

type User struct {
	Generic
	Name        string
	DateOfBirth string
	Address     string
	PrimaryRole string
	Email       string
	Password    string
	Active      *bool
}
