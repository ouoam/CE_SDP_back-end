package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Transcript struct {
	ID			null.Int	`json:"id" dont:"cu"`
	TourID		null.Int	`json:"tour_id" dont:"u"`
	UserID		null.Int	`json:"user_id" dont:"u"`
	File		null.String	`json:"file" dont:"u"`
	Confirm		null.Int	`json:"confirm"`
	IsCancel	null.Int	`json:"is_cancel"`
	Time		null.Time	`json:"time" dont:"cu"`
}

// todo time dont get when create

func (transcript *Transcript)SetID(id int64) {
	transcript.ID.SetValid(id)
}

func (transcript *Transcript) GetDB() error {
	err := db.GetData(transcript.ID.Int64, transcript)
	return err
}

func (transcript *Transcript) AddDB() error {
	if id, err := db.AddData(transcript); err != nil {
		return err
	} else {
		transcript.ID.SetValid(id)
	}
	return nil
}

func (transcript *Transcript) UpdateDB() error {
	err := db.UpdateDate(transcript.ID.Int64, transcript)
	return err
}