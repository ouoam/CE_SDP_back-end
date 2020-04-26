package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Review struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Time	null.Time	`json:"time" dont:"c"` //todo update time when update
}

func (review *Review) SetID(id int64)  {
	review.Tour.SetValid(id)
}

func (review *Review) GetDB() error {
	//err := db.GetData(review.ID.Int64, review)
	//return err

	return nil
}

func (review *Review) AddDB() error {
	return db.AddData(review)
}

func (review *Review) UpdateDB() error {
	return db.UpdateDate(review)
}

func (review *Review) ListDB() ([]interface{}, error) {
	results, err := db.ListData(review)
	return results, err
}
