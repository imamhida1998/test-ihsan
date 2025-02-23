package repository

import (
	"github.com/gin-gonic/gin"
	"test-ihsan/lib/db"
	"test-ihsan/model"
)

type bankRepository struct{}

func NewRepositoryBank() BankRepository {
	return &bankRepository{}

}

func (b *bankRepository) CreateBank(ctx *gin.Context, params *model.Bank) error {
	return db.PostgreSQL.Create(&params).Error
}

func (b *bankRepository) GetBank(ctx *gin.Context) ([]model.Bank, error) {
	var banks []model.Bank
	err := db.PostgreSQL.Find(&banks).Error
	if err != nil {
		return nil, err
	}
	return banks, nil

}
