package model

import (
	"gopkg.in/guregu/null.v3"
)

type Favorite struct {
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
}

func (favorite *Favorite) SetID(id int64)  {
	favorite.Tour.SetValid(id)
}