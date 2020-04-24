package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Review struct {
	ID		null.Int	`json:"id" dont:"cu"`
	UserID	null.Int	`json:"user_id" dont:"u"`
	TourID	null.Int	`json:"tour_id" dont:"u"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Score	null.Int	`json:"score"`
	Time	null.Time	`json:"time" dont:"c"`
}

//todo update time with now

func (review *Review)SetID(id int64) {
	review.ID.SetValid(id)
}

func (review *Review) GetDB() error {
	err := db.GetData(review.ID.Int64, review)
	return err
}

func (review *Review) AddDB() error {
	if id, err := db.AddData(review); err != nil {
		return err
	} else {
		review.ID.SetValid(id)
	}
	return nil
}

func (review *Review) UpdateDB() error {
	err := db.UpdateDate(review.ID.Int64, review)
	return err
}

func (review *Review) ListDB() ([]interface{}, error) {
	results, err := db.ListData(review)
	return results, err
}
