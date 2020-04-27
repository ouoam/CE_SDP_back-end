package model

import (
	"gopkg.in/guregu/null.v3"
)

type MessageListMe struct {
	Contact null.Int	`json:"contact" dont:"cud"`
	Me		null.Bool	`json:"me" dont:"cud"`
	Message	null.String	`json:"message" dont:"cud"`
	Time	null.Time	`json:"time" dont:"cud"`
}