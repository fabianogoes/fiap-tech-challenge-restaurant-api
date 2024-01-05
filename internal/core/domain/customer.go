package domain

import "time"

type Customer struct {
	ID        uint
	Name      string
	Email     string
	CPF       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(nome string, email string, cpf string) (*Customer, error) {
	return &Customer{
		Name:  nome,
		Email: email,
		CPF:   cpf,
	}, nil
}
