package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"test-ihsan/lib/logger"
	"test-ihsan/model"
	"test-ihsan/service/controller"
	"test-ihsan/service/usecase"
)

func Route(
	router *gin.Engine,
	Nasabah *usecase.UsecaseNasabah,
	auth *usecase.Auths) {
	handler := controller.NewControllerNasabah(Nasabah)

	api := router.Group("/api")

	api.POST("/daftar", handler.CreateNasabah)
	api.POST("/login", handler.Login)

	nasabah := api.Group("/nasabah")
	nasabah.Use(authMiddleware(*auth, *Nasabah))
	nasabah.POST("/tabung", handler.Nabung)
	nasabah.POST("/tarik", handler.Tarik)
	nasabah.GET("/saldo/:no_rekening", handler.Ceksaldo)
}

func authMiddleware(authService usecase.Auths, userService usecase.UsecaseNasabah) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		var tokenString string
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			logger.Log.Errorf("token validation error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			logger.Log.Error("token invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		nik := (claim["nik"].(string))

		user, err := userService.GetDetailNasabah(c, nik)
		if err != nil {
			logger.Log.Errorf("user detail error: %s", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		data := model.Nasabah{
			Id:           user.Id,
			Nik:          user.Nik,
			NoHp:         user.NoHp,
			Saldo:        user.Saldo,
			IdBank:       user.IdBank,
			Password:     user.Password,
			NoRekening:   user.NoRekening,
			PetugasRekam: user.PetugasRekam,
			TanggalRekam: user.TanggalRekam}

		c.Set("CurrentUser", data)

	}
}
