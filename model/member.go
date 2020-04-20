package model

import (
	"../db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"reflect"
	"regexp"
	"strconv"
)

type Member struct {
	ID           int		`json:"id" dontUpdate:""`
	Name         null.String`json:"name"`
	Surname      null.String`json:"surname"`
	Username     string		`json:"username" dontUpdate:""`
	Password     null.String`json:"password"`
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

	statement := `INSERT INTO public.member (name, surname, username, password, id_card, email, bank_account, address) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.DB.QueryRow(statement, member.Name, member.Surname, member.Username, member.Password, member.IdCard,
		member.Email, member.BankAccount, member.Address).Scan(&member.ID)

	_ = member.Password.UnmarshalText([]byte(""))
	return err
}

func UpdateMember(member *Member, id int) error {
	if err := preMember(member, false) ; err != nil {
		return err
	}

	var val []interface{}
	var sql string
	var count = 1

	stv := reflect.ValueOf(member).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		if _, have := fieldType.Tag.Lookup("dontUpdate"); have == true {
			continue
		}
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}

		valid := false
		v := field.Interface()

		switch v := v.(type) {
		case null.String:
			if v.Valid {
				valid = true
			}
		case null.Int:
			if v.Valid {
				valid = true
			}
		}
		if valid {
			sql += fieldType.Tag.Get("json") + " = $" + strconv.Itoa(count) + ", "
			count++
			val = append(val, v)
		}
	}

	if count == 1 {
		return errors.New("No data to update")
	}
	sql = sql[:len(sql) - 2]

	statement := "UPDATE public.member SET " + sql + " WHERE id=$" + strconv.Itoa(count) + ` 
RETURNING id, name, surname, username, id_card, email, verification, bank_account, address`
	val = append(val, id)

	err := db.DB.QueryRow(statement, val...).Scan(&member.ID, &member.Name, &member.Surname, &member.Username, &member.IdCard,
		&member.Email, &member.Verification, &member.BankAccount, &member.Address)
	if err != nil {
		return err
	}

	_ = member.Password.UnmarshalText([]byte(""))
	return nil
}