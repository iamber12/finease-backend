package models

type SupportTicket struct {
	Generic
	UserUUID    string
	Subject     string
	Description string
	Status      string
}
