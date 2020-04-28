package model

import (
	"../db"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type WithPostGet interface {
	PostGet() error
}

type WithPreChange interface {
	PreChange(isNew bool) error
}

type CanSearch interface {
	CanSearch() bool
}

func CheckValidAllPK(model interface{}) error {
	stv := reflect.ValueOf(model).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}
		v := field.Addr().Interface()
		valid := db.CheckValid(v)
		column := fieldType.Tag.Get("json")
		key := fieldType.Tag.Get("key")
		if strings.Contains(key, "p") && !valid {
			return errors.New("field " + column + " is invalid")
		}
	}
	return nil
}