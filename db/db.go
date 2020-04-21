package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq" //justify
	"gopkg.in/guregu/null.v3"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//DB a pointer to sql database
var DB *sql.DB

//Init postgresql db
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_DB"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Successfully connected to database!")
}

func UpdateDate(id int, data interface{}) error {
	var updateVal []interface{}
	var updateSQL string
	var returnVal []interface{}
	var returnSQL string
	var count = 1

	stv := reflect.ValueOf(data).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}
		v := field.Addr().Interface()
		val, have := fieldType.Tag.Lookup("dont")
		valid := false
		switch v := v.(type) {
		case *null.String:
			if v.Valid {
				valid = true
			}
		case *null.Int:
			if v.Valid {
				valid = true
			}
		}
		if valid && (!have || (have && strings.Contains(val, "u"))) {
			updateSQL += fieldType.Tag.Get("json") + " = $" + strconv.Itoa(count) + ", "
			count++
			updateVal = append(updateVal, v)
		}
		if !valid && (!have || (have && strings.Contains(val, "r"))) {
			returnSQL += fieldType.Tag.Get("json") + ", "
			returnVal = append(returnVal, v)
		}
	}

	if count == 1 {
		return errors.New("No data to update")
	}
	updateSQL = updateSQL[:len(updateSQL) - 2]
	returnSQL = returnSQL[:len(returnSQL) - 2]

	statement := "UPDATE public." + strings.ToLower(reflect.TypeOf(data).Elem().Name()) + " SET " + updateSQL
	statement += " WHERE id=$" + strconv.Itoa(count) + "RETURNING " + returnSQL
	updateVal = append(updateVal, id)

	err := DB.QueryRow(statement, updateVal...).Scan(returnVal...)
	return err
}