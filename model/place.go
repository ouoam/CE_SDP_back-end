package model

import (
	"gopkg.in/guregu/null.v3"
)

type Place struct {
	ID		null.Int	`json:"id" dont:"cu" key:"p"`
	Name	null.String	`json:"name"`
	Pic		null.String	`json:"pic"`
	Lat		null.Float	`json:"lat"`
	Lon		null.Float	`json:"lon"`
}

func (place *Place) SetID(id int64)  {
	place.ID.SetValid(id)
}