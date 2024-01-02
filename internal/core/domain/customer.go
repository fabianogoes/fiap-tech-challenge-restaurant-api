package domain

import "time"

type Customer struct {
	ID        uint
	Nome      string
	Email     string
	CPF       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(nome string, email string, cpf string) (*Customer, error) {
	return &Customer{
		Nome:  nome,
		Email: email,
		CPF:   cpf,
	}, nil
}
