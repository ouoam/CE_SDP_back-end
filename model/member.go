package model

import (
	"../db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"regexp"
)

type Member struct {
	ID           int		`json:"id"`
	Name         string		`json:"name"`
	Surname      string		`json:"surname"`
	Username     string		`json:"username"`
	Password     string		`json:"password"`
	IdCard       null.Int	`json:"id_card"`
	Email        string		`json:"email"`
	Verification null.Int	`json:"verification"`
	BankAccount  null.Int	`json:"bank_account"`
	Address      null.String`json:"address"`
}

var (
	// from http://emailregex.com/
	emailRegexp = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
)

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
	// TODO	- check id card https://th.wikipedia.org/wiki/Thai_ID https://th.wikipedia.org/wiki/ISO_3166-2:TH

	if len(member.Password) < 6 || 50 < len(member.Password) {
		return errors.New("Password length")
	}

	if !emailRegexp.MatchString(member.Email) {
		return errors.New("Email format")
	}

	//hash a password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)
	member.Password = string(bytes)

	statement := `INSERT INTO public.member (name, surname, username, password, id_card, email, bank_account, address) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.DB.QueryRow(statement, member.Name, member.Surname, member.Username, member.Password, member.IdCard,
		member.Email, member.BankAccount, member.Address).Scan(&member.ID)

	member.Password = ""

	return err
}

