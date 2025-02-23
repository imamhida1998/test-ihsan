package model

import "time"

type Bank struct {
	Id           string     `json:"id" gorm:"type:uuid;"`
	NamaBank     string     `json:"nama_bank" gorm:"type:varchar(255);not null"`
	TanggalRekam time.Time  `json:"tanggalRekam" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	PetugasUbah  *string    `json:"petugasUbah,omitempty" gorm:"type:varchar(100);default:null"`
	TanggalUbah  *time.Time `json:"tanggalUbah,omitempty" gorm:"type:timestamp NULL"`
}

func (Bank) TableName() string {
	return "bank"
}
