package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
	"test-ihsan/config"
	"test-ihsan/delivery"
	"test-ihsan/lib/db"
	"test-ihsan/lib/logger"
	"test-ihsan/service/repository"
	"test-ihsan/service/usecase"
)

var log = logrus.New()

func main() {
	config := config.Config{}
	config.CatchError(config.InitEnv())
	logger.InitLogger()
	var wg sync.WaitGroup

	router := gin.New()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := db.InitDBMysQL(config)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}()
	wg.Wait()

	if config.Get("MIGRATE") == "true" {
		db.Migrate()
	}
	router.Use(logger.LoggerMiddleware())
	router.Use(gin.Recovery())

	repoNasabah := repository.NewRepositoryNasabah()
	repoBank := repository.NewRepositoryBank()

	usecaseAUth := usecase.NewJWTService()

	usecasesNasabah := usecase.NewUsecaseNasabah(&repoNasabah, &usecaseAUth, &repoBank)
	delivery.Route(router, &usecasesNasabah, &usecaseAUth)

	router.Run("0.0.0.0:8081")

}
