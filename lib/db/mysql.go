package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"log"
	"test-ihsan/lib/logger"

	"test-ihsan/config"
)

var (
	PostgreSQL *gorm.DB
)

func InitDBMysQL(config config.Config) error {
	err := config.InitEnv()
	paramsDB := config.GetDBConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "invalid_config",
			"status": err.Error(),
		}).Error("Env init failed")
		return err
	}

	dataSourceName := "host=" + paramsDB.Host + " user=" + paramsDB.Username + " password=" + paramsDB.Password + " dbname=" + paramsDB.DBName + " port=" + paramsDB.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	PostgreSQL, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "invalid_config",
			"status": err.Error(),
		}).Error("Gagal koneksi db")
		return err
	}

	err = PostgreSQL.DB().Ping()
	if err != nil {

		log.Fatalf("failed to ping DB: %v", err)
	}

	logger.Log.Infof("Success koneksi db connected")
	return nil
}
