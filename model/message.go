package model

import (
	"gopkg.in/guregu/null.v3"
)

type Message struct {
	From	null.Int	`json:"from" dont:"ud" key:"p"`
	To		null.Int	`json:"to" dont:"ud" key:"p"`
	Time	null.Time	`json:"time" dont:"cud" key:"p"`
	Message	null.String	`json:"message" dont:"ud"`
}

type MessageListMe struct {
	Contact null.Int	`json:"contact" dont:"cud"`
	Me		null.Bool	`json:"me" dont:"cud"`
	Message	null.String	`json:"message" dont:"cud"`
	Time	null.Time	`json:"time" dont:"cud"`
	Name	null.String	`json:"name" dont:"cud"`
	Surname	null.String	`json:"surname" dont:"cud"`
	Pic		null.String	`json:"pic" dont:"cud"`
}

type MessageWithMe struct {
	Me		null.Bool	`json:"me" dont:"cud"`
	Message	null.String	`json:"message" dont:"cud"`
	Time	null.Time	`json:"time" dont:"cud"`
}