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

//todo update time with now

func (message *Message) GetDB() error {
	//err := db.GetData(message.ID.Int64, message)
	//return err

	return nil
}

func (message *Message) AddDB() error {
	//if id, err := db.AddData(message); err != nil {
	//	return err
	//} else {
	//	message.ID.SetValid(id)
	//}
	//return nil

	return nil
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
