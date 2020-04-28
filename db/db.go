package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
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

func CheckValid(v interface{}) bool {
	var valid bool
	switch v := v.(type) {
	case *null.String:
		valid = v.Valid
	case *null.Int:
		valid = v.Valid
	case *null.Time:
		valid = v.Valid
	case *null.Float:
		valid = v.Valid
	case *null.Bool:
		valid = v.Valid
	}
	return valid
}

func AddData(data interface{}) error {
	var insertVal []interface{}
	var insertSqlList string
	var insertSqlVal string
	var returnVal []interface{}
	var returnSql string
	var count = 1

	stv := reflect.ValueOf(data).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}
		v := field.Addr().Interface()
		dont := fieldType.Tag.Get("dont")
		valid := CheckValid(v)
		column := fieldType.Tag.Get("json")

		// change column of reserved word
		switch column {
		case "user":
			column = "\"user\""
		case "from":
			column = "\"from\""
		case "to":
			column = "\"to\""
		}

		if valid && !strings.Contains(dont, "c") {
			insertSqlList += column + ", "
			insertVal = append(insertVal, v)
			insertSqlVal += "$" + strconv.Itoa(count) + ", "
			count++
		} else {
			returnSql += column + ", "
			returnVal = append(returnVal, v)
		}
	}

	if count == 1 {
		return errors.New("no data to create")
	}
	insertSqlList = insertSqlList[:len(insertSqlList)-2]
	insertSqlVal = insertSqlVal[:len(insertSqlVal)-2]

	statement := "INSERT INTO public." + strings.ToLower(reflect.TypeOf(data).Elem().Name())
	statement += " (" + insertSqlList + ") VALUES (" + insertSqlVal + ")"
	if len(returnSql) != 0 {
		statement += " RETURNING " + returnSql[:len(returnSql)-2]
		return db.QueryRow(statement, insertVal...).Scan(returnVal...)
	}
	_, err := db.Exec(statement, insertVal...)
	return err
}

func UpdateDate(data interface{}) error {
	var updateVal []interface{}
	var updateSQL []string
	var whereVal []interface{}
	var whereSQL []string
	var returnVal []interface{}
	var returnSQL []string
	var count = 1

	stv := reflect.ValueOf(data).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}
		v := field.Addr().Interface()
		dont := fieldType.Tag.Get("dont")
		valid := CheckValid(v)
		column := fieldType.Tag.Get("json")
		key := fieldType.Tag.Get("key")

		// change column of reserved word
		switch column {
		case "user":
			column = "\"user\""
		case "from":
			column = "\"from\""
		case "to":
			column = "\"to\""
		}

		if key == "p" {
			if valid {
				whereSQL = append(whereSQL, column)
				whereVal = append(whereVal, v)
			} else {
				return errors.New("require key invalid")
			}
		} else if valid && !strings.Contains(dont, "u") {
			updateSQL = append(updateSQL, column + " = $" + strconv.Itoa(count))
			count++
			updateVal = append(updateVal, v)
		} else {
			returnSQL = append(returnSQL, column)
			returnVal = append(returnVal, v)
		}
	}

	if count == 1 {
		return errors.New("No data to update")
	}
	if len(whereSQL) == 0 {
		return errors.New("no where statement")
	}
	for i := range whereSQL{
		whereSQL[i] += " = $" + strconv.Itoa(count)
		count++
	}

	statement := "UPDATE public." + strings.ToLower(reflect.TypeOf(data).Elem().Name())
	statement += " SET " + strings.Join(updateSQL, ", ")
	statement += " WHERE " + strings.Join(whereSQL, " AND ")
	updateVal = append(updateVal, whereVal...)

	if len(returnSQL) != 0 {
		statement += " RETURNING " + strings.Join(returnSQL, ", ")
		return db.QueryRow(statement, updateVal...).Scan(returnVal...)
	}

	_, err := db.Exec(statement, updateVal...)
	return err
}

func ListData(data interface{}, params... interface{}) ([]interface{}, error) { // todo filter don't read data
	var argsList []interface{}
	var whereSQL []string
	var count = 1
	var fromParams []string

	for _, param := range params {
		fromParams = append(fromParams, "$" + strconv.Itoa(count))
		count++
		argsList = append(argsList, param)
	}

	stv := reflect.ValueOf(data).Elem()
	for i := 0; i < stv.NumField(); i++ {
		fieldType := stv.Type().Field(i)
		field := stv.Field(i)
		if !field.CanInterface() {
			continue
		}
		v := field.Addr().Interface()
		valid := CheckValid(v)
		column := fieldType.Tag.Get("json")
		switch column {
		case "user":
			column = "\"user\""
		case "from":
			column = "\"from\""
		case "to":
			column = "\"to\""
		}
		if valid {
			whereSQL = append(whereSQL, column + " = $"+strconv.Itoa(count))
			count++
			argsList = append(argsList, v)
		}
	}

	statement := "SELECT * FROM public." + strings.ToLower(reflect.TypeOf(data).Elem().Name())
	if len(fromParams) != 0 {
		statement += "(" + strings.Join(fromParams, ", ") + ")"
	}
	if len(whereSQL) != 0 {
		statement += " WHERE " + strings.Join(whereSQL, " AND ")
	}

	rows, err := db.Query(statement, argsList...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		result := reflect.New(stv.Type()).Interface()
		var returnVal []interface{}
		stv := reflect.ValueOf(result).Elem()
		for i := 0; i < stv.NumField(); i++ {
			field := stv.Field(i)
			if !field.CanInterface() {
				continue
			}
			v := field.Addr().Interface()
			switch field.Kind() {
			case reflect.Slice, reflect.Array:
				returnVal = append(returnVal, pq.Array(v))
			default:
				returnVal = append(returnVal, v)
			}
		}
		err := rows.Scan(returnVal...)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
