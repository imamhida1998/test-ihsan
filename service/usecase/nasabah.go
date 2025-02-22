package usecase

import (
	"test-ihsan/model/request"
	"test-ihsan/model/response"
)

type UsecaseNasabah interface {
	DaftarNasabah(params *request.Daftar) error
	Login(params *request.Login) (*response.GeneralResponse, error)
}
