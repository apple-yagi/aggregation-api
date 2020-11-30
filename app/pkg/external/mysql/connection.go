package mysql

import (
	"aggregation-mod/pkg/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(0.0.0.0:3306)/hoge?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&gateway.Experiment{}) {
		if err := db.Table("experiments").CreateTable(&gateway.Experiment{}).Error; err != nil {
			panic(err)
		}
	}

	if !db.HasTable(&gateway.Result{}) {
		if err := db.Table("results").CreateTable(&gateway.Result{}).Error; err != nil {
			panic(err)
		}
	}

	return db
}

func CloseConn() {
	db.Close()
}
