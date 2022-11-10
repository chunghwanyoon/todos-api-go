package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	logger "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewDB(settings Settings) (*sql.DB, error) {
	dbConfig := mysql.Config{
		Net:                  "tcp",
		User:                 settings.DB_USERNAME,
		Passwd:               settings.DB_PASSWORD,
		Addr:                 settings.DB_HOST,
		DBName:               settings.DB_DATABASE,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		logger.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	boil.SetDB(db)

	return db, nil
}
