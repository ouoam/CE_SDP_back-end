package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Message struct {
	From	null.Int	`json:"from" dont:"ud" key:"p"`
	To		null.Int	`json:"to" dont:"ud" key:"p"`
	Time	null.String	`json:"time" dont:"cud" key:"p"`
	Message	null.Int	`json:"message" dont:"ud"`
}

func (message *Message) SetID(id int64)  {
	message.To.SetValid(id)
}

func (message *Message) GetDB() error {
	//err := db.GetData(message.ID.Int64, message)
	//return err

	return nil
}

func (message *Message) AddDB() error {
	return db.AddData(message)
}

func (message *Message) UpdateDB() error {
	//err := db.UpdateDate(message.ID.Int64, message)
	//return err

	return nil
}

func (message *Message) ListDB() ([]interface{}, error) {
	results, err := db.ListData(message)
	return results, err
}
