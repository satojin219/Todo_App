package models

import (
	"config/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"

)

func init(){
	Db,err = sql.Open(config.Config.SQLDriver,config.Config.DbName)
	if err != nil{
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`,tableNameUser)

		_,err = Db.Exec(cmdU)
		if err != nil{
			log.Println(err)
		}

}