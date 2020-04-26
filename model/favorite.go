package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Favorite struct {
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
}

func (favorite *Favorite) SetID(id int64)  {
	favorite.Tour.SetValid(id)
}

func (favorite *Favorite) GetDB() error {
	//err := db.GetData(favorite.ID.Int64, favorite)
	//return err

	return nil
}

func (favorite *Favorite) AddDB() error {
	return db.AddData(favorite)
}

func (favorite *Favorite) UpdateDB() error {
	return db.UpdateDate(favorite)
}

func (favorite *Favorite) ListDB() ([]interface{}, error) {
	results, err := db.ListData(favorite)
	return results, err
}
