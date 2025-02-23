package repository

import "test-ihsan/model"

type RepositoryNasabah interface {
	CreateNasabah(params *model.Nasabah) error
	GetDetailNasabahByNIK(Nik string) (*model.Nasabah, error)
	CheckDataNasabahByNik(Nik string) (bool, error)
	CheckDataNasabahByNoHpAndPassword(noHp, password string) (*model.Nasabah, error)
	Nabung(nasabah *model.Nasabah) error
}
