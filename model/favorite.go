package model

import (
	"gopkg.in/guregu/null.v3"
)

type Favorite struct {
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
}

type FavoriteWithUser struct {
	Tour	null.Int	`json:"tour"`
	Name	null.String	`json:"name"`
}