package model

import (
	"gopkg.in/guregu/null.v3"
)

type List struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	Seq		null.Int	`json:"seq" dont:"u" key:"p"`
	Place	null.Int	`json:"place"`
}

type ListWithTour struct {
	ID		null.Int	`json:"id"`
	Name	null.String	`json:"name"`
	Pic		null.String	`json:"pic"`
	Lat		null.Float	`json:"lat"`
	Lon		null.Float	`json:"lon"`
}

type ListUpdate struct {
	ListUpdate	null.Int	`json:"list_update"`
}

type ListUpdateBody struct {
	Place	[]int64	`json:"place"`
}
