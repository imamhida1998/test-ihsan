package repository

import (
	"github.com/gin-gonic/gin"
	"test-ihsan/model"
)

type BankRepository interface {
	CreateBank(ctx *gin.Context, params *model.Bank) error
	GetBank(ctx *gin.Context) ([]model.Bank, error)
}
