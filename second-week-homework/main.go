package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func initUserTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS User;")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE User(Name text);")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err != nil {
		return err
	}
	return nil
}

/**
  dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层
  答：我认为如果是ErrNoRows的话，抛或不抛都行，如果抛则直接return error就行交给上层处理
      如果不抛，则进行兜底，返回空数组即可
*/
func main() {
	db, _ := sql.Open("sqlite3", "go-test.db")
	defer func() { _ = db.Close() }()
	err := initUserTable(db)
	if err != nil {
		fmt.Printf("初始化失败:%v", err)
		return
	}
	query, err := db.Query("select name from User")
	if err != nil {
		fmt.Printf("查询失败:%v", err)
		return
	}
	for query.Next() {
		var name string
		err := query.Scan(&name)
		if err != nil {
			fmt.Printf("扫描失败:%v", err)
			return
		}
		fmt.Printf("name:%s\n", name)
	}
}
