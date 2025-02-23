package db

import (
	"fmt"
	"test-ihsan/lib/logger"
	"test-ihsan/model"
)

func Migrate() {
	if exist := PostgreSQL.HasTable("nasabah"); !exist {
		fmt.Println("migrate table nasabah")
		err := PostgreSQL.CreateTable(&model.Nasabah{})
		if err != nil {
			logger.Log.Info("migrate table nasabah failed:", err.Error)
			return
		}
		fmt.Println("success migrate table nasabah")
	}

	if exist := PostgreSQL.HasTable("bank"); !exist {
		fmt.Println("migrate table bank")
		err := PostgreSQL.CreateTable(&model.Bank{})
		if err != nil {
			logger.Log.Info("migrate table bank failed:", err.Error)
			return
		}
		logger.Log.Info("migrate table bank")
	}
}
