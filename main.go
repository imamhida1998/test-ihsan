package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
	"test-ihsan/config"
	"test-ihsan/lib/db"
	"test-ihsan/service/controller"
	"test-ihsan/service/repository"
	"test-ihsan/service/usecase"
	"time"
)

var log = logrus.New()

func main() {
	config := config.Config{}
	config.CatchError(config.InitEnv())
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

	db.Migrate()

	repoNasabah := repository.NewRepositoryNasabah()

	usecaseAUth := usecase.NewJWTService()
	usecasesNasabah := usecase.NewUsecaseNasabah(&repoNasabah, &usecaseAUth)

	controllerNasabah := controller.NewControllerNasabah(&usecasesNasabah)

	controllerNasabah.Routes(router)

	router.Use(gin.Recovery())

	router.Use(LoggerMiddleware())
	router.Run("127.0.0.1:8080")

}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())
		c.Set("request_id", requestID)

		c.Next()

		// Ambil informasi setelah request selesai
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// Log request
		log.WithFields(logrus.Fields{
			"request_id": requestID,
			"status":     statusCode,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"latency":    duration,
			"ip":         clientIP,
		}).Info("Request processed")
	}
}
