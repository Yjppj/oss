package main

import (
	"fmt"
	"go.uber.org/zap"
	"sso/dao/mysql"
	"sso/initialize"
	"sso/initialize/logger"
	"sso/routers"
)

func main() {
	err := initialize.Init()
	if err != nil {
		fmt.Printf("init settings failed ,err:%v\n", err)
		zap.L().Error(fmt.Sprintf("init settings failed ,err:%v\n", err))
		return
	}

	zap.L().Debug("viper init success...")
	//初始化Zap
	if err := logger.Init(initialize.Conf.LogConfig, initialize.Conf.Mode); err != nil {
		fmt.Printf("init logger failed ,err:%v\n", err)
		zap.L().Error(fmt.Sprintf("init logger failed ,err:%v\n", err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("zap init success...")
	if err := mysql.Init(initialize.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed ,err:%v\n", err)
		zap.L().Error(fmt.Sprintf("init mysql failed ,err:%v\n", err))
		return
	}

	r := routers.Setup(initialize.Conf.Mode)
	r.Run(fmt.Sprintf(":%d", initialize.Conf.App.Port))
}
