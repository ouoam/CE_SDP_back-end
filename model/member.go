package model

import (
	"../db"
)

type Member struct {
	ID           int	`json:"id"`
	Name         string	`json:"name"`
	Surname      string	`json:"surname"`
	Username     string	`json:"username"`
	Password     string	`json:"password"`
	IdCard       uint64	`json:"id_card"`
	Email        string	`json:"email"`
	Verification uint8	`json:"verification"`
	BankAccount  uint64	`json:"bank_account"`
	Address      string	`json:"address"`
}

func GetMember(id int) (*Member, error) {
	member := new(Member)
	statement := `SELECT id, name, surname, username, id_card, email, verification, bank_account, address 
FROM public.member WHERE id = $1`
	err := db.DB.QueryRow(statement, id).Scan(&member.ID, &member.Name, &member.Surname, &member.Username, &member.IdCard,
		&member.Email, &member.Verification, &member.BankAccount, &member.Address)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func AddMember(member *Member) error {
	// TODO	- encrypt password
	//		- check email
	//		- check id card

	statement := `INSERT INTO public.member (name, surname, username, password, id_card, email, bank_account, address) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.DB.QueryRow(statement, member.Name, member.Surname, member.Username, member.Password, member.IdCard,
		member.Email, member.BankAccount, member.Address).Scan(&member.ID)

	return err
}

