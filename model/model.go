package model

type WithID interface {
	GetDB() error
	AddDB() error
	UpdateDB() error
	ListDB() ([]interface{}, error)
}
