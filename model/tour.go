package model

import (
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Tour struct {
	ID			null.Int	`json:"id" dont:"cu" key:"p"`
	Owner		null.Int	`json:"owner" dont:"u" key:"p"` // use key p for check is owner
	Name		null.String	`json:"name"`
	Description	null.String	`json:"description"`
	Category	null.String	`json:"category"`
	MaxMember	null.Int	`json:"max_member"`
	FirstDay	null.Time	`json:"first_day"`
	LastDay		null.Time	`json:"last_day"`
	Price		null.Int	`json:"price"`
	Status		null.Int	`json:"status"`
	Pic			null.String	`json:"pic"`
}

func (tour *Tour) PreChange(isNew bool) error {
	if !tour.FirstDay.Valid || !tour.LastDay.Valid {
		return errors.New("first day and last day is requires")
	}
	if tour.FirstDay.Time.After(tour.LastDay.Time) {
		return errors.New("first day is after last day")
	}
	if tour.FirstDay.Time.Before(time.Now()) {
		return errors.New("first day is before now")
	}
	return nil
}

type TourDetailSearch struct {
	ID			null.Int	`json:"id" key:"p"`
	Owner		null.Int	`json:"owner"`
	Name		null.String	`json:"name"`
	Description	null.String	`json:"description"`
	Category  	null.String `json:"category"`
	MaxMember 	null.Int    `json:"max_member"`
	FirstDay  	null.Time   `json:"first_day"`
	LastDay   	null.Time   `json:"last_day"`
	Price     	null.Int    `json:"price"`
	Status    	null.Int    `json:"status"`
	Member    	null.Int    `json:"member"`
	Confirm   	null.Int    `json:"confirm"`
	Ratting   	null.Float  `json:"ratting"`
	Favorite  	null.Int    `json:"favorite"`
	GName     	null.String `json:"g_name"`
	GSurname  	null.String `json:"g_surname"`
	BankAccount	null.Int    `json:"bank_account"`
	BankName  	null.String `json:"bank_name"`
	List	  	[]string	`json:"list"`
}

func (tour *TourDetailSearch) CanSearch() bool {
	return true
}