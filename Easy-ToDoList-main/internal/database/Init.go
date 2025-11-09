package database

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"todolist/config"
	"todolist/internal/model"
	"todolist/pkg/logUtil"
)

var db *gorm.DB

var modelList = []interface{}{
	model.User{},
	model.ToDo{},
}

func Init() *gorm.DB {
	sqliteConfig := config.Get().Sqlite
	sqliteDB, err := gorm.Open(sqlite.New(sqlite.Config{DriverName: sqliteConfig.Driver, DSN: sqliteConfig.DataSourceName}))
	if err != nil {
		logUtil.Logger.Error("数据库连接失败", zap.Error(err))
	}
	db = sqliteDB

	GenerateTable()

	return db
}

func GenerateTable() {
	err := db.AutoMigrate(modelList...)
	if err != nil {
		logUtil.Logger.Error("数据表创建失败", zap.Error(err))
	}
}

func GetDB() *gorm.DB {
	return db
}
