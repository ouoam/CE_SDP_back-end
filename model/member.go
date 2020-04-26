package model

import (
	"../db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"regexp"
)

type Member struct {
	ID          null.Int    `json:"id" dont:"cu" key:"p"`
	Name        null.String `json:"name"`
	Surname     null.String `json:"surname"`
	Username    null.String `json:"username" dont:"u"`
	Password    null.String `json:"password" dont:"r"`
	IdCard      null.Int    `json:"id_card"`
	Email       null.String `json:"email"`
	BankAccount null.Int    `json:"bank_account"`
	Address     null.String `json:"address"`
	Verify		null.Bool	`json:"verify"`
}

var (
	// from http://emailregex.com/
	emailRegexp = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
)

func (member *Member) preMember(isNew bool) error {
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

func (member *Member) GetDB() error {
	if err := db.GetData(member.ID.Int64, member); err != nil {
		return err
	}

	_ = member.Password.UnmarshalText([]byte(""))
	return nil
}

func (member *Member) AddDB() error {
	if err := member.preMember(true); err != nil {
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

func (member *Member) UpdateDB() error {
	if err := member.preMember(false); err != nil {
		return err
	}

	if err := db.UpdateDate(member.ID.Int64, member); err != nil {
		return err
	}

	_ = member.Password.UnmarshalText([]byte(""))
	return nil
}

func (member *Member) ListDB() ([]interface{}, error) {
	results, err := db.ListData(member)
	if err != nil {
		return nil, err
	}
	members := make([]interface{}, len(results))
	for i, result := range results {
		temp := result.(*Member)
		_ = temp.Password.UnmarshalText([]byte(""))
		members[i] = temp
	}
	return members, nil
}
