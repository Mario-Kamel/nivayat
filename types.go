package main

type Makhdoum struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
	Address   string `json: "address"`
}

func NewMakhdoum(firstName, lastName, address string) *Makhdoum {
	return &Makhdoum{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
	}
}
