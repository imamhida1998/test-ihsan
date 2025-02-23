package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log = logrus.New()

func InitLogger() {

	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)

		Log.SetLevel(logrus.InfoLevel)
	}
}

func Info(c *gin.Context, function, msg string) {
	start := time.Now()
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())

	duration := time.Since(start)
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()

	entry := Log.WithFields(logrus.Fields{
		"request_id": requestID,
		"function":   function,
		"status":     statusCode,
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"latency":    fmt.Sprintf("%v", duration),
		"ip":         clientIP,
		"message":    msg,
		"time":       time.Now().Format(time.RFC3339), // Tambahkan timestamp
	})

	entry.WithField("msg", "Request processed").Info("")

}

func Error(c *gin.Context, function, msg string) {
	start := time.Now()
	requestID := fmt.Sprintf("%d", time.Now().UnixNano())

	duration := time.Since(start)
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()

	entry := Log.WithFields(logrus.Fields{
		"request_id": requestID,
		"function":   function,
		"status":     statusCode,
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"latency":    fmt.Sprintf("%v", duration),
		"ip":         clientIP,
		"message":    msg,
		"time":       time.Now().Format(time.RFC3339), // Tambahkan timestamp
	})

	entry.WithField("msg", "Request failed").Error("")

}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := fmt.Sprintf("%d", time.Now().UnixNano()) // Generate request ID unik
		c.Set("request_id", requestID)

		c.Next() // Jalankan request

		// Ambil informasi setelah request selesai
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		entry := Log.WithFields(logrus.Fields{
			"request_id": requestID,
			"status":     statusCode,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"latency":    fmt.Sprintf("%v", duration),
			"ip":         clientIP,
			"time":       time.Now().Format(time.RFC3339), // Tambahkan timestamp
		})

		if statusCode >= 500 {
			entry.WithField("msg", "Request failed").Error("")
		} else {
			entry.WithField("msg", "Request processed").Info("")
		}
	}
}
