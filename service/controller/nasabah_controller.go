package controller

import (
	"github.com/gin-gonic/gin"
	"test-ihsan/model/request"
	"test-ihsan/service/usecase"
)

type NasabahController struct {
	NasabahService usecase.UsecaseNasabah
}

func NewControllerNasabah(nasabah *usecase.UsecaseNasabah) NasabahController {
	return NasabahController{
		NasabahService: *nasabah,
	}
}

func (c *NasabahController) Routes(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/daftar", c.CreateNasabah)
	api.POST("/login", c.login)

}

func (c *NasabahController) CreateNasabah(ctx *gin.Context) {
	var params request.Daftar

	err := ctx.ShouldBind(&params)
	if err != nil {
		//ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
		//	StatusCode:    http.StatusBadRequest,
		//	StatusMessage: "Permintaan ditolak",
		//})
		return
	}

	err = c.NasabahService.DaftarNasabah(&params)
	if err != nil {
		return
	}

	return

}

func (c *NasabahController) login(ctx *gin.Context) {
	var params request.Login

	err := ctx.ShouldBind(&params)
	if err != nil {
		//ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
		//	StatusCode:    http.StatusBadRequest,
		//	StatusMessage: "Permintaan ditolak",
		//})
		return
	}

	res, err := c.NasabahService.Login(&params)
	if err != nil {
		return
	}

	ctx.JSON(200, res)

}
