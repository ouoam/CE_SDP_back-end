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
var db *sql.DB

//Init postgresql db
func Init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_DB"))
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}

func AddData(data interface{}) (int64, error) {
	var insertVal []interface{}
	var insertSQL string
	var insertSqlVal string
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
			valid = v.Valid
		case *null.Int:
			valid = v.Valid
		case *null.Time:
			valid = v.Valid
		}
		if valid && (!have || (have && !strings.Contains(val, "c"))) {
			insertSQL += fieldType.Tag.Get("json") + ", "
			insertVal = append(insertVal, v)
			insertSqlVal += "$" + strconv.Itoa(count) + ", "
			count++
		}
	}

	if count == 1 {
		return 0, errors.New("No data to update")
	}
	insertSQL = insertSQL[:len(insertSQL) - 2]
	insertSqlVal = insertSqlVal[:len(insertSqlVal) - 2]

	statement := "INSERT INTO public." + strings.ToLower(reflect.TypeOf(data).Elem().Name()) + " (" + insertSQL + ") "
	statement += "VALUES (" + insertSqlVal + ") RETURNING id"

	var id int64
	err := db.QueryRow(statement, insertVal...).Scan(&id)

	return id, err
}

func GetData(id int64, data interface{}) error {
	var getSQL string
	var returnVal []interface{}

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
			valid = v.Valid
		case *null.Int:
			valid = v.Valid
		case *null.Time:
			valid = v.Valid
		}
		if !valid && (!have || (have && !strings.Contains(val, "r"))) {
			getSQL += fieldType.Tag.Get("json") + ", "
			returnVal = append(returnVal, v)
		}
	}

	getSQL = getSQL[:len(getSQL) - 2]

	statement := "SELECT " + getSQL + " FROM public." + strings.ToLower(reflect.TypeOf(data).Elem().Name()) + " WHERE id = $1"
	err := db.QueryRow(statement, id).Scan(returnVal...)

	return err
}

func UpdateDate(id int64, data interface{}) error {
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
			valid = v.Valid
		case *null.Int:
			valid = v.Valid
		case *null.Time:
			valid = v.Valid
		}
		if valid && (!have || (have && !strings.Contains(val, "u"))) {
			updateSQL += fieldType.Tag.Get("json") + " = $" + strconv.Itoa(count) + ", "
			count++
			updateVal = append(updateVal, v)
		}
		if (!valid && (!have || (have && !strings.Contains(val, "r")))) ||
			(valid && (!have || (have && strings.Contains(val, "u")))) {
			returnSQL += fieldType.Tag.Get("json") + ", "
			returnVal = append(returnVal, v)
		}
	}

	if count == 1 {
		return errors.New("No data to update")
	}

	statement := "UPDATE public." + strings.ToLower(reflect.TypeOf(data).Elem().Name())
	statement += " SET " + updateSQL[:len(updateSQL) - 2]
	statement += " WHERE id=$" + strconv.Itoa(count)
	updateVal = append(updateVal, id)

	if len(returnSQL) != 0 {
		statement += " RETURNING " + returnSQL[:len(returnSQL)-2]
		err := db.QueryRow(statement, updateVal...).Scan(returnVal...)
		return err
	}

	_, err := db.Exec(statement, updateVal...)
	return err
}