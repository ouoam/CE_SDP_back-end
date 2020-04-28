package db

import (
	"gopkg.in/guregu/null.v3"
)

type Filter struct {
	Page 	null.Int	`json:"page"`
	Limit	null.Int	`json:"limit"`
	Offset	null.Int	`json:"offset"`
	Order	null.String	`json:"order"`
	Search	null.String	`json:"search"`
	Desc	bool		`json:"desc"`
	Valid	bool
}

func (filter *Filter) PreUse() {
	if !filter.Limit.Valid {
		filter.Limit.SetValid(30)
	}
	if filter.Order.Valid {
		filter.Order.SetValid(EscapeReserveWord(filter.Order.String))
	}
	if filter.Page.Valid && !filter.Offset.Valid{
		filter.Offset.SetValid(filter.Page.Int64 * filter.Limit.Int64)
	}
}