package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solerwell/go-learn/logger"
	"github.com/solerwell/go-learn/util"
	"os"
)

var log2 = logger.GetStdLogger()

func main() {
	path := getSqlitePath()
	fmt.Println("sqlite datasource file path : ", path)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log2.Fatal("get sqlite3 db error: ", err.Error())
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log2.Fatal(err)
	}
	rows, _ := tx.Query("select * from product where id = ?", 1)
	defer rows.Close()
	if rows.Next() {
		fmt.Println("we found a record in the database!")
	} else {
		fmt.Println("no record is found by the id - ", 1)
	}
}

func getSqlitePath() string {
	rootPath := util.GetProjectRootPath()
	sep := os.PathSeparator
	return fmt.Sprintf("%s%cyanyue_spider%cdata%cyanyue.db", rootPath, sep, sep, sep)
}
