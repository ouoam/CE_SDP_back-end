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

func (Review *Review)SetID(id int64) {
	Review.ID.SetValid(id)
}

func (Review *Review) GetDB() error {
	err := db.GetData(Review.ID.Int64, Review)
	return err
}

func (Review *Review) AddDB() error {
	if id, err := db.AddData(Review); err != nil {
		return err
	} else {
		Review.ID.SetValid(id)
	}
	return nil
}

func (Review *Review) UpdateDB() error {
	err := db.UpdateDate(Review.ID.Int64, Review)
	return err
}