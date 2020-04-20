package model

import (
	"../db"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v3"
	"regexp"
	"strconv"
)

type Member struct {
	ID           int		`json:"id"`
	Name         null.String`json:"name"`
	Surname      null.String`json:"surname"`
	Username     string		`json:"username"`
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
	var i = 1

	if member.Name.Valid {
		sql += "name = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Name.String)
	}

	if member.Surname.Valid {
		sql += "surname = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Surname.String)
	}

	if member.Password.Valid {
		sql += "password = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Password.String)
	}

	if member.IdCard.Valid {
		sql += "id_card = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.IdCard.Int64)
	}

	if member.Email.Valid {
		sql += "email = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Email.String)
	}

	if member.Verification.Valid {
		sql += "verification = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Verification.Int64)
	}

	if member.BankAccount.Valid {
		sql += "bank_account = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.BankAccount.Int64)
	}

	if member.Address.Valid {
		sql += "address = $" + strconv.Itoa(i) + ", "
		i++
		val = append(val, member.Address.String)
	}

	if i == 1 {
		return errors.New("No data to update")
	}
	sql = sql[:len(sql) - 2]

	statement := "UPDATE public.member SET " + sql + " WHERE id=$" + strconv.Itoa(i) + ` 
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