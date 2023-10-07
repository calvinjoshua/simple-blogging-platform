package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type blog struct {
	Id     int
	Author string
	Blog   string
}

func dbConnection() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "", 5432, "", "", "")

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	return db, nil

}

func checkIdExist(id int) (bool, error) {

	var conn *sql.DB
	var err error
	var res *sql.Rows

	conn, err = dbConnection()
	if err != nil {
		return false, err
	}

	res, err = conn.Query("SELECT * FROM TABLENAME WHERE ID=$1", id)

	defer res.Close()

	if res.Next() {
		return true, nil
	}

	return false, errors.New("Id does not exist")

}

func insertBlogData(id int, blog string, author string) (bool, error) {
	var conn *sql.DB
	var err error
	conn, err = dbConnection()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	_, err = conn.Exec("INSERT INTO TABLENAME(ID, BLOG,AUTHOR) VALUE ($1, $2)", id, blog, author)

	if err != nil {
		return false, err
	}

	return true, nil
}

func retriveBlog(id int) (interface{}, error) {

	var conn *sql.DB
	var err error
	var res *sql.Rows
	var idValidation bool

	idValidation, err = checkIdExist(id)

	if !idValidation {
		return nil, err

	}

	conn, err = dbConnection()
	if err != nil {
		return nil, err
	}

	res, err = conn.Query("SELECT * FROM TABLENAME WHERE ID= $1", id)

	if err != nil {
		return nil, err
	}

	defer res.Close()
	var _blog blog

	if res.Next() {

		err = res.Scan(_blog)

		if err != nil {
			return nil, err
		}

	}

	return _blog, nil
}

func getAllblogs() (interface{}, error) {

	var _blog []blog
	var conn *sql.DB
	var err error

	var res *sql.Rows

	conn, err = dbConnection()
	if err != nil {
		return nil, err
	}

	res, err = conn.Query("SELECT * FROM TABLENAME")

	defer res.Close()

	for res.Next() {

		var temp blog
		err = res.Scan(temp)
		if err != nil {
			return nil, err
		}

		_blog = append(_blog, temp)

	}

	return _blog, nil
}

func _deleteBlog(id int) error {

	var conn *sql.DB
	var err error

	var idValidation bool

	idValidation, err = checkIdExist(id)

	if !idValidation {
		return err

	}

	conn, err = dbConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec("DELETE FROM TABLENAME WHERE ID = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func _updateBlog(id int, blog string) (bool, error) {

	var conn *sql.DB
	var err error
	var idValidation bool

	idValidation, err = checkIdExist(id)

	if !idValidation {
		return false, err

	}

	conn, err = dbConnection()
	if err != nil {
		return false, err
	}

	_, err = conn.Exec("UPDATE TABLE SET BLOG=$2 WHERE ID=$2", blog, id)

	if err != nil {
		return false, err
	}

	return true, nil

}
