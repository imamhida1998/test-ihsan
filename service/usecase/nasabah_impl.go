package usecase

import (
	"github.com/google/uuid"
	"test-ihsan/helpers"
	"test-ihsan/model"
	"test-ihsan/model/request"
	"test-ihsan/model/response"

	//"test-ihsan/model/response"
	"test-ihsan/service/repository"
)

type nasabahUsecase struct {
	RepoNasabah repository.RepositoryNasabah
	Auth        Auths
}

func NewUsecaseNasabah(nasabah *repository.RepositoryNasabah, auth *Auths) UsecaseNasabah {
	return &nasabahUsecase{RepoNasabah: *nasabah,
		Auth: *auth}
}

func (u *nasabahUsecase) DaftarNasabah(params *request.Daftar) error {

	statusCheck, err := u.RepoNasabah.CheckDataNasabahByNik(params.Nik)
	if err != nil {
		return err
	}
	if !statusCheck {
		id, _ := uuid.NewV7()
		reqDaftar := &model.Nasabah{
			Id:         id.String(),
			Nik:        params.Nik,
			NoHp:       params.NoHp,
			NoRekening: helpers.GenerateRekening(10),
			Saldo:      0,
			Password:   helpers.EncryptedHash(params.Password),
		}

		err = u.RepoNasabah.CreateNasabah(reqDaftar)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *nasabahUsecase) Nabung() {

}

func (u *nasabahUsecase) Login(params *request.Login) (*response.GeneralResponse, error) {

	status, err := u.RepoNasabah.CheckDataNasabahByNoHpAndPassword(params.NoHp, helpers.EncryptedHash(params.Password))
	if err != nil {
		return nil, err
	}

	token, err := u.Auth.GenerateTokenJWT(status.Nik)
	if err != nil {
		return nil, err
	}

	return &response.GeneralResponse{
		Code: 200,
		Msg:  "success",
		Data: map[string]string{
			"token": token,
		},
	}, nil

}
