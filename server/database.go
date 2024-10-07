package server

import (
	"database/sql"
	"fmt"
)

type DbOption uint8

const (
	LOAD DbOption = iota
	INSERT
	UPDATE
	DELETE
)

type DbConf struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Schema   string
}

func DbServer(district *DbConf) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", district.Username, district.Password, district.Host, district.Port, district.Schema)
	db, e := sql.Open("mysql", dsn)
	if e != nil {
		return nil, e
	}

	if e := db.Ping(); e != nil {
		return nil, e
	}

	return db, nil
}