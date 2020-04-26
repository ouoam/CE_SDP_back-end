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

//todo update time with now

func (review *Review) GetDB() error {
	//err := db.GetData(review.ID.Int64, review)
	//return err

	return nil
}

func (review *Review) AddDB() error {
	//if id, err := db.AddData(review); err != nil {
	//	return err
	//} else {
	//	review.ID.SetValid(id)
	//}
	//return nil

	return nil
}

func (review *Review) UpdateDB() error {
	//err := db.UpdateDate(review.ID.Int64, review)
	//return err

	return nil
}

func (review *Review) ListDB() ([]interface{}, error) {
	results, err := db.ListData(review)
	return results, err
}
