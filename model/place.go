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

func GetPlace(id int) (*Place, error) {
	place := new(Place)
	if err := db.GetData(id, place); err != nil {
		return nil, err
	}
	return place, nil
}

func AddPlace(place *Place) error {
	if id, err := db.AddData(place); err != nil {
		return err
	} else {
		place.ID.SetValid(id)
	}
	return nil
}

func UpdatePlace(place *Place, id int) error {
	if err := db.UpdateDate(id, place); err != nil {
		return err
	}
	return nil
}