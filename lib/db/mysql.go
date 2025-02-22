package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"

	"test-ihsan/config"
)

var (
	Mysql *gorm.DB
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
	dataSourceName := paramsDB.DBName + ":" + paramsDB.Password + "@tcp(" + paramsDB.Host + ":" + paramsDB.Port + ")/" + paramsDB.DBName + "?parseTime=true"

	Mysql, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"event":  "invalid_config",
			"status": err.Error(),
		}).Error("Gagal koneksi db")
		return err
	}

	err = Mysql.DB().Ping()
	if err != nil {

		log.Fatalf("failed to ping DB: %v", err)
	}

	log.Println("Database connected successfully!")
	return nil
}
