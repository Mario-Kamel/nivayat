package main

import "time"

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName, address string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		CreatedAt: time.Now().UTC(),
	}
}
