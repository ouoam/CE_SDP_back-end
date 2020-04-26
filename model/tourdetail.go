package model

import (
	"../db"
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v3"
)

type TourDetail struct {
	ID			null.Int	`json:"id" key:"p"`
	Owner		null.Int	`json:"owner"`
	Name		null.String	`json:"name"`
	Description	null.String	`json:"description"`
	Category	null.String	`json:"category"`
	MaxMember	null.Int	`json:"max_member"`
	FirstDay	null.Time	`json:"first_day"`
	LastDay		null.Time	`json:"last_day"`
	Price		null.Int	`json:"price"`
	Status		null.Int	`json:"status"`
	Member		null.Int	`json:"member"`
	Confirm		null.Int	`json:"confirm"`
	Ratting		null.Float	`json:"ratting"`
}

// todo check first day will before last day

func (tour *TourDetail) SetID(id int64)  {
	tour.ID.SetValid(id)
}

func (tour *TourDetail) GetDB() error {
	err := db.GetData(tour.ID.Int64, tour)
	return err
}

func (tour *TourDetail) AddDB() error {
	return errors.New("this struct can not be add")
}

func (tour *TourDetail) UpdateDB() error {
	return errors.New("this struct can not be update")
}

func (tour *TourDetail) ListDB() ([]interface{}, error) {
	results, err := db.ListData(tour)
	return results, err
}