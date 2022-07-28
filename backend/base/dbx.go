package base

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"time"
)

//dbx 数据库实例
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

type DbxDatabaseStarter struct {
}

func (s *DbxDatabaseStarter) Setup() {

	settings := dbx.Settings{
		DriverName: "mysql",
		Host:            "106.55.46.203:3306",
		Database:        "db_town_guide",
		User:            "root",
		Password:        "dlyj520",
		ConnMaxLifetime: time.Minute * 30,
		MaxOpenConns:    10,
		MaxIdleConns:    2,
		Options: map[string]string{
			"charset":   "utf8",
			"parseTime": "true",
			"loc":       "Local",
		}}
	log.Info("mysql.conn url:", settings.ShortDataSourceName())
	db, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	log.Info(db.Ping())
	database = db
}
