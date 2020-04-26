package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Favorite struct {
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
}

func (favorite *Favorite) GetDB() error {
	//err := db.GetData(favorite.ID.Int64, favorite)
	//return err

	return nil
}

func (favorite *Favorite) AddDB() error {
	//if id, err := db.AddData(favorite); err != nil {
	//	return err
	//} else {
	//	favorite.ID.SetValid(id)
	//}
	//return nil

	return nil
}

func (favorite *Favorite) UpdateDB() error {
	//err := db.UpdateDate(favorite.ID.Int64, favorite)
	//return err

	return nil
}

func (favorite *Favorite) ListDB() ([]interface{}, error) {
	results, err := db.ListData(favorite)
	return results, err
}
