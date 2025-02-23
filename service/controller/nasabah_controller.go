package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"test-ihsan/helpers/constant"
	"test-ihsan/lib/logger"
	"test-ihsan/model/request"
	"test-ihsan/model/response"
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

func (c *NasabahController) CreateNasabah(ctx *gin.Context) {
	var params request.Daftar

	err := ctx.ShouldBind(&params)
	if err != nil {
		logger.Error(ctx, constant.FC_CREATE_NASABAH, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	convertResp, _ := json.Marshal(&params)
	logger.Info(ctx, constant.FC_LOGIN, string(convertResp))
	res, err := c.NasabahService.DaftarNasabah(ctx, &params)
	if err != nil {
		logger.Error(ctx, constant.FC_CREATE_NASABAH, err.Error())
		resp := &response.GeneralResponse{
			Code: 400,
			Msg:  "gagal daftar nasabah",
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := &response.GeneralResponse{
		Code: 200,
		Msg:  "Success",
		Data: res,
	}

	ctx.JSON(http.StatusOK, resp)

	return

}

func (c *NasabahController) Login(ctx *gin.Context) {
	var params request.Login

	err := ctx.ShouldBind(&params)
	if err != nil {
		logger.Error(ctx, constant.FC_LOGIN, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}
	convertResp, _ := json.Marshal(&params)
	logger.Info(ctx, constant.FC_LOGIN, string(convertResp))
	res, err := c.NasabahService.Login(ctx, &params)
	if err != nil {
		logger.Error(ctx, constant.FC_LOGIN, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	ctx.JSON(200, res)

}

func (c *NasabahController) Nabung(ctx *gin.Context) {
	var params request.Tabung

	err := ctx.ShouldBind(&params)
	if err != nil {
		logger.Error(ctx, constant.FC_NABUNG, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	res, err := c.NasabahService.Nabung(ctx, &params)
	if err != nil {
		logger.Error(ctx, constant.FC_NABUNG, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	ctx.JSON(200, res)

}

func (c *NasabahController) Tarik(ctx *gin.Context) {
	var params request.Tabung

	err := ctx.ShouldBind(&params)
	if err != nil {
		logger.Error(ctx, constant.FC_TARIK, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	res, err := c.NasabahService.TarikSaldo(ctx, &params)
	if err != nil {
		logger.Error(ctx, constant.FC_TARIK, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	ctx.JSON(200, res)

}

func (c *NasabahController) Ceksaldo(ctx *gin.Context) {

	var params request.CheckSaldo
	convertResp, _ := json.Marshal(&params)
	logger.Log.Info("CreateNasabah", string(convertResp))
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		logger.Error(ctx, constant.FC_CEK_SALDO, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	res, err := c.NasabahService.CekSaldo(ctx, &params)
	if err != nil {
		logger.Error(ctx, constant.FC_CEK_SALDO, err.Error())
		ctx.JSON(http.StatusBadRequest, response.GeneralResponse{
			Code: http.StatusBadRequest,
			Msg:  "Permintaan ditolak",
		})
		return
	}

	ctx.JSON(200, res)

}
