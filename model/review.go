package model

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

type Review struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	User	null.Int	`json:"user" dont:"u" key:"p"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Time	null.Time	`json:"time" dont:"c"`
}

func (review *Review) PreChange(isNew bool) error {
	if !isNew {
		review.Time.SetValid(time.Now())
	}
	return nil
}

type ReviewWithUser struct {
	Tour	null.Int	`json:"tour" key:"p"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Time	null.Time	`json:"time"`
	Name	null.String	`json:"name"`
}

type ReviewWithTour struct {
	User	null.Int	`json:"user" key:"p"`
	Comment	null.String	`json:"comment"`
	Ratting	null.Int	`json:"ratting"`
	Time	null.Time	`json:"time"`
	Name	null.String	`json:"name"`
	Surname	null.String	`json:"surname"`
}