package repo

import (
	"context"
	"log"

	"github.com/dotdak/exchange-system/dao"
	"gorm.io/gorm"
)

type WagerRepo interface {
	Create(ctx context.Context, wager *dao.Wager) (*dao.Wager, error)
	Update(ctx context.Context, wager *dao.Wager) (*dao.Wager, error)
	List(ctx context.Context, offset, limit uint) ([]*dao.Wager, error)
	Get(ctx context.Context, id uint32) (*dao.Wager, error)
}

type WagerRepoImpl struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewWagerRepo(
	db *gorm.DB,
	logger *log.Logger,
) WagerRepo {
	db.AutoMigrate(&dao.Wager{})
	return &WagerRepoImpl{
		logger: logger,
		db:     db,
	}
}

func (r *WagerRepoImpl) Create(ctx context.Context, wager *dao.Wager) (*dao.Wager, error) {
	if err := r.db.Create(wager).Error; err != nil {
		return nil, err
	}

	return wager, nil
}

func (r *WagerRepoImpl) Update(ctx context.Context, wager *dao.Wager) (*dao.Wager, error) {
	if err := r.db.Model(&wager).Select(
		"CurrentSellingPrice", "PercentageSold", "AmountSold",
	).Updates(&wager).Error; err != nil {
		return nil, err
	}

	return wager, nil
}

func (r *WagerRepoImpl) List(ctx context.Context, offset, limit uint) ([]*dao.Wager, error) {
	var wagers []*dao.Wager
	if err := r.db.Model(wagers).Limit(int(limit)).Offset(int(offset)).Find(&wagers).Error; err != nil {
		return nil, err
	}

	return wagers, nil
}

func (r *WagerRepoImpl) Get(ctx context.Context, id uint32) (*dao.Wager, error) {
	var wager dao.Wager
	if err := r.db.First(&wager, id).Error; err != nil {
		return nil, err
	}

	return &wager, nil
}
