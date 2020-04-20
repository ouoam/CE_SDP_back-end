package model

import (
	"../db"
)

type Place struct {
	ID		int     `json:"id"`
	Name	string  `json:"name"`
	Pic		string  `json:"pic"`
	Geo		string	`json:"geo"`
}

func GetPlace(id int) (*Place, error) {
	place := new(Place)
	statement := `SELECT id, name, pic, geo FROM public.place WHERE id = $1`
	err := db.DB.QueryRow(statement, id).Scan(&place.ID, &place.Name, &place.Pic, &place.Geo)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func AddPlace(place *Place) error {
	statement := `INSERT INTO public.place(name, pic, geo) VALUES($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(statement, place.Name, place.Pic, place.Geo).Scan(&place.ID)

	return err
}

