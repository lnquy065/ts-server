package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"ts-server/configs"
)

var Pgdb gorm.DB

func PgConnect() {
	fmt.Println(configs.ConnectionString)
	pgdb, err := gorm.Open(configs.Dialect, configs.ConnectionString)

	if err != nil {
		fmt.Println("DB error: " + err.Error())
	} else {
		pgdb.SingularTable(true)
		pgdb.LogMode(true)

		Pgdb = *pgdb
	}
}
