package usecase

import (
	"github.com/gin-gonic/gin"
	"test-ihsan/model"
	"test-ihsan/model/request"
	"test-ihsan/model/response"
)

type UsecaseNasabah interface {
	DaftarNasabah(ctx *gin.Context, params *request.Daftar) (*model.Nasabah, error)
	Login(ctx *gin.Context, params *request.Login) (*response.GeneralResponse, error)
	GetDetailNasabah(ctx *gin.Context, nik string) (*model.Nasabah, error)
	Nabung(ctx *gin.Context, params *request.Tabung) (*model.Nasabah, error)
	TarikSaldo(ctx *gin.Context, params *request.Tabung) (*model.Nasabah, error)
	CekSaldo(ctx *gin.Context, params *request.CheckSaldo) (*model.Nasabah, error)
}
