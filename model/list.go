package model

import (
	"gopkg.in/guregu/null.v3"
)

type List struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	Seq		null.Int	`json:"seq" dont:"u" key:"p"`
	Place	null.Int	`json:"place"`
}

func (list *List) SetID(id int64)  {
	list.Tour.SetValid(id)
}