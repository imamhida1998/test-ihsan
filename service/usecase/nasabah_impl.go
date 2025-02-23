package usecase

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"test-ihsan/helpers"
	"test-ihsan/helpers/constant"
	"test-ihsan/lib/logger"
	"test-ihsan/model"
	"test-ihsan/model/request"
	"test-ihsan/model/response"

	//"test-ihsan/model/response"
	"test-ihsan/service/repository"
)

type nasabahUsecase struct {
	RepoNasabah repository.RepositoryNasabah
	Bank        repository.BankRepository
	Auth        Auths
}

func NewUsecaseNasabah(nasabah *repository.RepositoryNasabah, auth *Auths, bank *repository.BankRepository) UsecaseNasabah {
	return &nasabahUsecase{RepoNasabah: *nasabah,
		Auth: *auth,
		Bank: *bank}
}

func (u *nasabahUsecase) DaftarNasabah(ctx *gin.Context, params *request.Daftar) (*model.Nasabah, error) {

	statusCheck, err := u.RepoNasabah.CheckDataNasabahByNik(params.Nik)
	if err != nil {
		return nil, err
	}

	getBank, err := u.Bank.GetBank(ctx)
	if err != nil {
		return nil, err
	}

	if getBank == nil {
		return nil, errors.New("bank tidak tersedia")
	}

	if statusCheck {
		return nil, errors.New("nasabah sudah terdaftar")
	}
	id, _ := uuid.NewV7()
	reqDaftar := &model.Nasabah{
		Id:           id.String(),
		Nik:          params.Nik,
		NoHp:         params.NoHp,
		IdBank:       getBank[0].Id,
		NoRekening:   helpers.GenerateRekening(10),
		PetugasRekam: "ADMIN",
		Password:     helpers.EncryptedHash(params.Password),
	}

	err = u.RepoNasabah.CreateNasabah(reqDaftar)
	if err != nil {
		return nil, err
	}
	return reqDaftar, nil

	return nil, errors.New("nasabah telah terdaftar")

}

func (u *nasabahUsecase) GetDetailNasabah(ctx *gin.Context, nik string) (*model.Nasabah, error) {

	res, err := u.RepoNasabah.GetDetailNasabahByNIK(nik)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *nasabahUsecase) Nabung(ctx *gin.Context, params *request.Tabung) (*model.Nasabah, error) {
	currentUser := ctx.MustGet("CurrentUser").(model.Nasabah)

	result, err := u.RepoNasabah.GetDetailNasabahByNIK(currentUser.Nik)
	if err != nil {
		return nil, err
	}

	if result.NoRekening != params.NoRekening {
		return nil, errors.New("nomor rekening tidak sesuai")
	}
	saldo, err := strconv.ParseFloat(params.Nominal, 64)
	if err != nil {
		return nil, err

	}
	result.Saldo += saldo

	err = u.RepoNasabah.Nabung(result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (u *nasabahUsecase) Login(ctx *gin.Context, params *request.Login) (*response.GeneralResponse, error) {

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

func (u *nasabahUsecase) TarikSaldo(ctx *gin.Context, params *request.Tabung) (*model.Nasabah, error) {
	currentUser := ctx.MustGet("CurrentUser").(model.Nasabah)

	result, err := u.RepoNasabah.GetDetailNasabahByNIK(currentUser.Nik)
	if err != nil {

		return nil, err
	}

	if result.NoRekening != params.NoRekening {
		return nil, errors.New("nomor rekening tidak sesuai")
	}

	saldo, err := strconv.ParseFloat(params.Nominal, 64)
	if err != nil {
		return nil, err

	}

	if saldo > result.Saldo {
		return nil, errors.New("Saldo anda tidak cukup")
	}
	logger.Info(ctx, constant.FU_TARIK_SALDO, fmt.Sprintf("saldo nasabah %f", result.Saldo))
	logger.Info(ctx, constant.FU_TARIK_SALDO, fmt.Sprintf("nasabah tarik saldo sebesar %f", saldo))

	result.Saldo -= saldo

	err = u.RepoNasabah.Nabung(result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (u *nasabahUsecase) CekSaldo(ctx *gin.Context, params *request.CheckSaldo) (*model.Nasabah, error) {

	currentUser := ctx.MustGet("CurrentUser").(model.Nasabah)

	if currentUser.NoRekening != params.NoRekening {
		return nil, errors.New("nomor rekening tidak sesuai")
	}

	res, err := u.RepoNasabah.GetDetailNasabahByNIK(currentUser.Nik)
	if err != nil {
		return nil, err
	}
	return res, nil

}
