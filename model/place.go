package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Place struct {
	ID		null.Int	`json:"id" dont:"cu"`
	Name	null.String	`json:"name"`
	Pic		null.String	`json:"pic"`
	Geo		null.String	`json:"geo"`
}

func (place *Place) GetDB() error {
	err := db.GetData(place.ID.Int64, place)
	return err
}

func (place *Place) AddDB() error {
	if id, err := db.AddData(place); err != nil {
		return err
	} else {
		place.ID.SetValid(id)
	}
	return nil
}

func (place *Place) UpdateDB() error {
	err := db.UpdateDate(place.ID.Int64, place)
	return err
}