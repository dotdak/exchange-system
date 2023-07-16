package repo

import (
	"context"

	"github.com/dotdak/exchange-system/dao"
	"gorm.io/gorm"
)

type BuyRepo interface {
	Create(ctx context.Context, buy *dao.Buy) (*dao.Buy, error)
}

type BuyRepoImpl struct {
	db *gorm.DB
}

func NewBuyRepo(
	db *gorm.DB,
) BuyRepo {
	db.AutoMigrate(&dao.Buy{})
	return &BuyRepoImpl{
		db: db,
	}
}

func (r *BuyRepoImpl) Create(ctx context.Context, buy *dao.Buy) (*dao.Buy, error) {
	if err := r.db.Create(buy).Error; err != nil {
		return nil, err
	}

	return buy, nil
}
