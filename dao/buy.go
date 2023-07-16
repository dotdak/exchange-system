package dao

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Buy struct {
	BaseModel
	WagerID     uint32    `gorm:"column:wager_id" validate:"required"`
	BuyingPrice float64   `gorm:"column:buying_price" validate:"required"`
	BoughtAt    time.Time `gorm:"column:bought_at" validate:"required"`
}

func (b *Buy) TableName() string {
	return "buys"
}

// BeforeSave validate Wager model
func (b *Buy) BeforeSave(tx *gorm.DB) error {
	err := validate.Struct(b)
	if err != nil {
		return fmt.Errorf("can't save invalid wager: %w", err)
	}
	return nil
}
