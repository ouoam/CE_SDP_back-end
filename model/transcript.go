package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Transcript struct {
	Tour		null.Int	`json:"tour" dont:"u" key:"p"`
	User		null.Int	`json:"user" dont:"u" key:"p"`
	File		null.String	`json:"file" dont:"u" key:"p"`
	Confirm		null.Int	`json:"confirm"`
	Time		null.Time	`json:"time" dont:"cu"`
}

func (transcript *Transcript) SetID(id int64)  {
	transcript.Tour.SetValid(id)
}

func (transcript *Transcript) GetDB() error {
	//err := db.GetData(transcript.ID.Int64, transcript)
	//return err

	return nil
}

func (transcript *Transcript) AddDB() error {
	return db.AddData(transcript)
}

func (transcript *Transcript) UpdateDB() error {
	return db.UpdateDate(transcript)
}

func (transcript *Transcript) ListDB() ([]interface{}, error) {
	results, err := db.ListData(transcript)
	return results, err
}