package mysql

import (

	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sso/global"
	"sso/initialize"
	"time"
)

func Init(cfg *initialize.MySQLConfig) (err error) {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		//DSN: "root:123456@tcp(121.43.119.224:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local",
		DSN: DSN, // 1. 连接信息
	}), &gorm.Config{ // 2. 选项
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true, //不用物理外键，使用逻辑外键
	})
	if err != nil {
		zap.L().Debug("数据库链接失败")
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10) //
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.GLOAB_DB = db
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

