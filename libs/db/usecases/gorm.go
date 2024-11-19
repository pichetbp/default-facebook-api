package usecases

import (
	"default-repo/helpers"
	logModels "default-repo/libs/log/models"
	logUsecases "default-repo/libs/log/usecases"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB GormConInterface

type GormConInterface interface {
	GetDB() *gorm.DB
}

type gormCon struct {
	db *gorm.DB
}

func NewGormCon() GormConInterface {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	gormConn, err := gorm.Open(sqlserver.Open(helpers.GetConnectionString()), gormConfig)
	if err != nil {
		logUsecases.LogEntry.OnError(logModels.LogError{
			Error: err.Error(),
		})
	}

	gormDB, err := gormConn.DB()
	if err != nil {
		logUsecases.LogEntry.OnError(logModels.LogError{
			Error: err.Error(),
		})
	}
	gormDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.maxLifeTime")) * time.Minute)
	gormDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	gormDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))

	GormDB := &gormCon{
		db: gormConn,
	}

	return GormDB
}

func (g *gormCon) GetDB() *gorm.DB {
	return g.db
}
