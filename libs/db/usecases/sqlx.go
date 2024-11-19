package usecases

import (
	"default-repo/helpers"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var SqlxDB SqlxConInterface

type SqlxConInterface interface {
	GetDB() *sqlx.DB
}

type sqlxCon struct {
	db *sqlx.DB
}

func NewSqlxCon() SqlxConInterface {
	sqlxConn, err := sqlx.Open(helpers.GetDriverString(), helpers.GetConnectionString())
	if err != nil {
		// add log
	}
	sqlxConn.SetConnMaxLifetime(time.Duration(viper.GetInt("db.maxLifeTime")) * time.Minute)
	sqlxConn.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlxConn.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))

	SqlxDB = &sqlxCon{
		db: sqlxConn,
	}

	return SqlxDB
}

func (s *sqlxCon) GetDB() *sqlx.DB {
	return s.db
}
