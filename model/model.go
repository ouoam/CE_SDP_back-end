package model

type WithID interface {
	SetID(id int64)
	GetDB() error
	AddDB() error
	UpdateDB() error
	ListDB() ([]interface{}, error)
}
