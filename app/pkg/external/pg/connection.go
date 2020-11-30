package pg

import (
	"aggregation-mod/pkg/adapter/gateway"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kelseyhightower/envconfig"
)

var db *gorm.DB

type Env struct {
	Sslmode  string `default:"disable"`
	Port     string `default:"5432"`
	Host     string `default:"localhost"`
	DBNAME   string `default:"test"`
	User     string `default:"postgres"`
	Password string `default:"password"`
}

func Connect() *gorm.DB {
	var err error
	var pgenv Env
	if err := envconfig.Process("pg", &pgenv); err != nil {
		panic(err)
	}

	db, err = gorm.Open("postgres", "host="+pgenv.Host+" port="+pgenv.Port+" user="+pgenv.User+" dbname="+pgenv.DBNAME+" password="+pgenv.Password+" sslmode="+pgenv.Sslmode)

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
