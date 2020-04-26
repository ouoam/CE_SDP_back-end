package model

import (
	"../db"
	"gopkg.in/guregu/null.v3"
)

type List struct {
	Tour	null.Int	`json:"tour" dont:"u" key:"p"`
	Seq		null.Int	`json:"seq" dont:"u" key:"p"`
	Place	null.Int	`json:"place"`
}

func (list *List) GetDB() error {
	//err := db.GetData(list.ID.Int64, list)
	//return err

	return nil
}

func (list *List) AddDB() error {
	//if id, err := db.AddData(list); err != nil {
	//	return err
	//} else {
	//	list.ID.SetValid(id)
	//}
	//return nil

	return nil
}

func (list *List) UpdateDB() error {
	//err := db.UpdateDate(list.ID.Int64, list)
	//return err

	return nil
}

func (list *List) ListDB() ([]interface{}, error) {
	results, err := db.ListData(list)
	return results, err
}
