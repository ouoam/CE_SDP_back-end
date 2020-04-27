package model

import (
	"gopkg.in/guregu/null.v3"
)

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
	List	  	[]string	`json:"list"`
}