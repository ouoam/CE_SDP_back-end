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
}

// todo check first day will before last day

func (tour *Tour) SetID(id int64)  {
	tour.ID.SetValid(id)
}
