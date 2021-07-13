package initialize

import (
	"gin-DevOps/config"
	"gin-DevOps/model"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func mysqlAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.Permission{},
	)
	if err != nil {
		config.GdoLog.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}

	config.GdoLog.Info("migrate table success")
}

func GormMysql() *gorm.DB {
	m := config.GdoConfig.Mysql
	mysqlConfig := ""
	for index, conf := range m.Config{
		if index == 0{
			mysqlConfig += conf
		}else {
			mysqlConfig += "&" + conf
		}
	}
	mysqlClient := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + mysqlConfig
	if db, err := gorm.Open(mysql2.Open(mysqlClient)); err != nil {
		config.GdoLog.Error("MySQL连接异常", zap.Any("err", err))
		return nil
	} else {
		//sqlDB, _ := db.DB()
		//sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		//sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		//mysqlAutoMigrate(db)
		return db
	}
}