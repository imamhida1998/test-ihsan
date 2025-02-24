package model

import (
	"time"
)

type Nasabah struct {
	Id           string     `json:"id" gorm:"type:uuid;"`
	Nik          string     `json:"nik" gorm:"type:varchar(100);unique"`
	NoHp         string     `json:"no_hp" gorm:"type:varchar(30)"`
	Saldo        float64    `json:"saldo" gorm:"type:DECIMAL(15,2);"`
	NoRekening   string     `json:"no_rekening" gorm:"type:varchar(30)"`
	IdBank       string     `json:"id_bank" gorm:"type:varchar(36);"`
	Password     string     `json:"password" gorm:"type:varchar(100)"`
	PetugasRekam string     `json:"petugasRekam" gorm:"type:varchar(100)"`
	TanggalRekam time.Time  `json:"tanggalRekam" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	PetugasUbah  *string    `json:"petugasUbah,omitempty" gorm:"type:varchar(100);default:null"`
	TanggalUbah  *time.Time `json:"tanggalUbah,omitempty" gorm:"type:timestamp NULL"`
}

func (Nasabah) TableName() string {
	return "nasabah"

}
