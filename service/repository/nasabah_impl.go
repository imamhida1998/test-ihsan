package repository

import (
	"errors"
	"test-ihsan/lib/db"
	"test-ihsan/model"
)

type Nasabah struct{}

func NewRepositoryNasabah() RepositoryNasabah {
	return &Nasabah{}
}

func (r *Nasabah) CreateNasabah(params *model.Nasabah) error {
	return db.Mysql.Create(&params).Error
}

func (r *Nasabah) GetDetailNasabahByNIK(Nik string) (*model.Nasabah, error) {
	var nasabah model.Nasabah
	err := db.Mysql.First(&nasabah, "NIK = ?", Nik).Error
	if err != nil {
		return nil, err
	}
	return &nasabah, nil

}

func (r *Nasabah) CheckDataNasabahByNik(Nik string) (bool, error) {
	var nasabah model.Nasabah
	var count int64
	db.Mysql.Model(&nasabah).Where("NIK = ?", Nik).Count(&count)

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *Nasabah) CheckDataNasabahByNoHpAndPassword(noHp, password string) (*model.Nasabah, error) {
	var nasabah model.Nasabah
	//var count int64
	err := db.Mysql.Model(&nasabah).Where(" no_hp = ? and password = ?", noHp, password).Find(&nasabah).Error
	if err != nil {
		return nil, err
	}
	if nasabah.Id != "" {
		return &nasabah, nil
	} else {
		return nil, errors.New("nasabah not found")
	}
}
