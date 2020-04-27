package model

import (
	"gopkg.in/guregu/null.v3"
)

type Review struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Time	null.Time	`json:"time" dont:"c"` //todo update time when update
}
