package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satoukick/webserver/config"
	logs "github.com/satoukick/webserver/log"
)

var db *gorm.DB

func init() {
	config.Init()
	initPostgres()
}

func initPostgres() {
	var err error
	pgconf := config.Conf.GetPGEnvString()
	db, err = gorm.Open("postgres", pgconf)
	if err != nil {
		logs.Fatal(err)
	}
	db.AutoMigrate(TodoModel{})
}
