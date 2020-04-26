package model

import (
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