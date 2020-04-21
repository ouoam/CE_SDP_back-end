package model

import (
	"../db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"regexp"
)

type Member struct {
	ID           null.Int	`json:"id" dont:"cu"`
	Name         null.String`json:"name"`
	Surname      null.String`json:"surname"`
	Username     null.String`json:"username" dont:"u"`
	Password     null.String`json:"password" dont:"r"`
	IdCard       null.Int	`json:"id_card"`
	Email        null.String`json:"email"`
	Verification null.Int	`json:"verification"`
	BankAccount  null.Int	`json:"bank_account"`
	Address      null.String`json:"address"`
}

var (
	// from http://emailregex.com/
	emailRegexp = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
)

func preMember(member *Member, isNew bool) error {
	// TODO	- check id card https://th.wikipedia.org/wiki/Thai_ID https://th.wikipedia.org/wiki/ISO_3166-2:TH
	//		- check username
	//		- check name and surname

	if isNew && !member.Password.Valid {
		return errors.New("no password")
	}

	if l := len(member.Password.String); l != 0 && (l < 6 || 50 < l) {
		return errors.New("Password length")
	}

	if isNew && !member.Email.Valid {
		return errors.New("no email")
	}

	if len(member.Email.String) != 0 && !emailRegexp.MatchString(member.Email.String) {
		return errors.New("Email format")
	}

	if member.Password.Valid {
		bytes, err := bcrypt.GenerateFromPassword([]byte(member.Password.String), bcrypt.DefaultCost)
		member.Password.String = string(bytes)

		return err
	}

	return nil
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

	_ = member.Password.UnmarshalText([]byte(""))
	return member, nil
}

func AddMember(member *Member) error {
	if err := preMember(member, true) ; err != nil {
		return err
	}

	if id, err := db.AddData(member); err != nil {
		return err
	} else {
		member.ID.SetValid(id)
	}

	_ = member.Password.UnmarshalText([]byte(""))
	return nil
}

func UpdateMember(member *Member, id int) error {
	if err := preMember(member, false) ; err != nil {
		return err
	}

	if err := db.UpdateDate(id, member); err != nil {
		return err
	}

	_ = member.Password.UnmarshalText([]byte(""))
	return nil
}