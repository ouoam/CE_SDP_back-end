package model

import (
	"../db"
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

func (place *Place) GetDB() error {
	err := db.GetData(place.ID.Int64, place)
	return err
}

func (place *Place) AddDB() error {
	return db.AddData(place)
}

func (place *Place) UpdateDB() error {
	err := db.UpdateDate(place.ID.Int64, place)
	return err
}

func (place *Place) ListDB() ([]interface{}, error) {
	results, err := db.ListData(place)
	return results, err
}
