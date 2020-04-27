package model

import (
	"gopkg.in/guregu/null.v3"
)

type Tour struct {
	ID			null.Int	`json:"id" dont:"cu" key:"p"`
	Owner		null.Int	`json:"owner" dont:"u"`
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

// todo check first day will before last day

type TourDetail struct {
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