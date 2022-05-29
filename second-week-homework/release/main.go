package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

/**
  定义错误码
*/
var notFoundCode = 40001
var systemErr = 50001

func Biz(sql string) (*sql.Rows, error) {
	rows, err := Dao(sql)
	if IsNoRow(err) {
		//根据业务来进行定义，如果可以允许为nil返回nil就行，如业务不允许，可以考虑转换为另一种错误，或直接返回错误响应
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return rows, err
}

func Dao(query string) (*sql.Rows, error) {
	db, _ := sql.Open("sqlite3", "go-test.db")
	defer func() { _ = db.Close() }()
	rows, err := db.Query(query)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("%d, sql:%s, not found", notFoundCode, query)
	} else if err != nil {
		return nil, fmt.Errorf("%d, sql:%s, %v", systemErr, query, err)
	}
	return rows, nil
}

func IsNoRow(err error) bool {
	return strings.HasPrefix(err.Error(), fmt.Sprintf("%d", notFoundCode))
}

func main() {
	biz, err := Biz("select * from User")
	if err != nil {
		log.Printf("query err %v", err)
		return
	}
	for biz.Next() {
		var value string
		err := biz.Scan(&value)
		if err != nil {
			log.Printf("scan err %v", err)
			return
		}
		fmt.Printf("result: %s \n", value)
	}
}
