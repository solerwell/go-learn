package yanyue_spider

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solerwell/go-learn/logger"
	"github.com/solerwell/go-learn/util"
	"github.com/solerwell/go-learn/yanyue_spider/yanyue"
	"os"
)

var log = logger.GetStdLogger()

func Write2Sqlite(pros []*yanyue.Product) {
	dsPath := getSqlitePath()
	db, err := sql.Open("sqlite3", dsPath)
	if err != nil {
		log.Fatal("get sqlite3 db error: ", err.Error())
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	selectSql := "select * from product where id = ?"
	insertSql := `insert into product(id, name, brand, barprice, packprice, commentnum, pingnum, comsore, scorewei, scorebao, scorejia) 
				  VALUES (?,?,?,?,?,?,?,?,?,?,?)`
	updateSql := `update product set name=?,brand=?, barprice=?, packprice=?, commentnum=?, pingnum=?, 
                   comsore=?, scorewei=?, scorebao=?, scorejia=? where id = ?`
	for i := 0; i < len(pros); i++ {
		pro := pros[i]
		rows, _ := tx.Query(selectSql, pro.Id)
		defer rows.Close()
		if rows.Next() {
			_, err := tx.Exec(insertSql, pro.Id, pro.Name, pro.Brand, pro.BarPrice, pro.PackPrice,
				pro.CommentNum, pro.PingNum, pro.ComSore, pro.ScoreWei, pro.ScoreBao, pro.ScoreJia)
			if err != nil {
				log.Errorf("insert cigarette error, id = %d,error: %s\n", pro.Id, err.Error())
			}
		} else {
			_, err := tx.Exec(updateSql, pro.Name, pro.Brand, pro.BarPrice, pro.PackPrice,
				pro.CommentNum, pro.PingNum, pro.ComSore, pro.ScoreWei, pro.ScoreBao, pro.ScoreJia, pro.Id)
			if err != nil {
				log.Errorf("update cigarette error, id = %d,error: %s\n", pro.Id, err.Error())
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Errorf("commit transaction error: %s\n", err.Error())
		_ := tx.Rollback()
	}
}

func Write2Csv(pro *yanyue.Product) {
	// todo
	// to be implemented
}

func getSqlitePath() string {
	rootPath := util.GetProjectRootPath()
	sep := os.PathSeparator
	return fmt.Sprintf("%s%cyanyue_spider%cdata%cyanyue.db", rootPath, sep, sep, sep)
}
