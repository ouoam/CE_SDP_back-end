package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type Tour struct {
	ID			null.Int	`json:"id" dont:"cu" key:"p"`
	Owner		null.Int	`json:"owner" dont:"u"`
	Name		null.String	`json:"name"`
	Description	null.String	`json:"description"`
	Category	null.String	`json:"category"`
	MaxMember	null.Int	`json:"max_member"`
	FirstDay	null.Time	`json:"first_day"`
	LastDay		null.Time	`json:"last_day"`
	Price		null.Int	`json:"price"`
	Status		null.Int	`json:"status"`
}

// todo check first day will before last day

func (tour *Tour) GetDB() error {
	err := db.GetData(tour.ID.Int64, tour)
	return err
}

func (tour *Tour) AddDB() error {
	if id, err := db.AddData(tour); err != nil {
		return err
	} else {
		tour.ID.SetValid(id)
	}
	return nil
}

func (tour *Tour) UpdateDB() error {
	err := db.UpdateDate(tour.ID.Int64, tour)
	return err
}

func (tour *Tour) ListDB() ([]interface{}, error) {
	results, err := db.ListData(tour)
	return results, err
}