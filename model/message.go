package model

import (
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