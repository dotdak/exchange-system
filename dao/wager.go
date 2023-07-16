package dao

import (
	"fmt"
	"time"

	v1 "github.com/dotdak/exchange-system/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
)

type Wager struct {
	BaseModel
	TotalWagerValue     uint32    `json:"total_wager_value" gorm:"column:total_wager_value" validate:"gt=0"`
	Odds                uint32    `json:"odds" gorm:"column:odds" validate:"gt=0"`
	SellingPercentage   uint32    `json:"selling_percentage" gorm:"column:selling_percentage" validate:"gte=1,lte=100"`
	SellingPrice        float64   `json:"selling_price" gorm:"column:selling_price" validate:"required"`
	CurrentSellingPrice float64   `json:"current_selling_price" gorm:"column:current_selling_price"`
	PercentageSold      *uint32   `json:"percentage_sold" gorm:"column:percentage_sold"`
	AmountSold          *float64  `json:"amount_sold" gorm:"column:amount_sold"`
	PlacedAt            time.Time `json:"placed_at" gorm:"column:placed_at"`
}

func (w *Wager) TableName() string {
	return "wagers"
}

func (w *Wager) FromProto(pb *v1.CreateWagerResponse) *Wager {
	v := &Wager{
		TotalWagerValue:     pb.TotalWagerValue,
		Odds:                pb.Odds,
		SellingPercentage:   pb.SellingPercentage,
		SellingPrice:        pb.SellingPrice,
		CurrentSellingPrice: pb.CurrentSellingPrice,
		PlacedAt:            pb.PlacedAt.AsTime(),
	}
	if pb.PercentageSold != nil {
		v.PercentageSold = &pb.PercentageSold.Value
	}
	if pb.AmountSold != nil {
		v.AmountSold = &pb.AmountSold.Value
	}

	*w = *v
	return v
}

func (w *Wager) ToProto() *v1.CreateWagerResponse {
	v := &v1.CreateWagerResponse{
		Id:                  uint32(w.ID),
		TotalWagerValue:     w.TotalWagerValue,
		Odds:                w.Odds,
		SellingPercentage:   w.SellingPercentage,
		SellingPrice:        w.SellingPrice,
		CurrentSellingPrice: w.CurrentSellingPrice,
		PlacedAt:            timestamppb.New(w.PlacedAt),
	}

	if w.AmountSold != nil {
		v.AmountSold = wrapperspb.Double(*w.AmountSold)
	}

	if w.PercentageSold != nil {
		v.PercentageSold = wrapperspb.UInt32(*w.PercentageSold)
	}

	return v
}

// BeforeSave validate Wager model
func (w *Wager) BeforeSave(tx *gorm.DB) error {
	err := validate.Struct(w)
	if err != nil {
		return fmt.Errorf("can't save invalid wager: %w", err)
	}
	return nil
}
