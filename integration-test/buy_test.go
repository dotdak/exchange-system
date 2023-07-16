//go:build integration

package integration

import (
	"context"
	"time"

	"github.com/dotdak/exchange-system/dao"
	v1 "github.com/dotdak/exchange-system/proto/v1"
)

func (p *MySqlRepositoryTestSuite) TestBuyFlow() {
	ctx := context.Background()
	p.Run("Given valid buy then should save", func() {
		// Given
		wager := &dao.Wager{
			BaseModel:           dao.BaseModel{ID: 1},
			TotalWagerValue:     123,
			Odds:                123,
			SellingPercentage:   10,
			SellingPrice:        123,
			CurrentSellingPrice: 123,
			PlacedAt:            time.Now(),
		}
		buy := &v1.BuyRequest{
			WagerId:     1,
			BuyingPrice: 100,
		}

		// When
		_, err := p.wagerRepo.Create(ctx, wager)
		// Then
		p.Assert().NoError(err)

		// When
		buyRes, err := p.handler.Buy(ctx, buy)
		// Then
		p.Assert().NoError(err)
		p.Assert().Equal(buyRes.Id, uint32(1))
		actualBuy := &dao.Buy{}
		err = p.db.First(actualBuy, buyRes.Id).Error
		p.Assert().NoError(err)
		p.Assert().Equal(float64(100), actualBuy.BuyingPrice)
	})

	p.Run("Given buy with no wager id then should not save", func() {
		// Given
		wager := &dao.Wager{
			BaseModel:           dao.BaseModel{ID: 0},
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        1,
			CurrentSellingPrice: 1,
			PlacedAt:            time.Now(),
		}
		buy := &v1.BuyRequest{
			WagerId:     0,
			BuyingPrice: 0.25,
		}

		// When
		_, err := p.wagerRepo.Create(ctx, wager)
		// Then
		p.Assert().NoError(err)
		// When
		_, err = p.handler.Buy(ctx, buy)
		// Then
		p.Assert().Error(err)
	})

	p.Run("Given buy with no price then should not save", func() {
		// Given
		wager := &dao.Wager{
			BaseModel:           dao.BaseModel{ID: 1},
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        1,
			CurrentSellingPrice: 1,
			PlacedAt:            time.Now(),
		}
		buy := &v1.BuyRequest{
			WagerId: 1,
		}

		// When
		_, err := p.wagerRepo.Create(ctx, wager)
		// Then
		p.Assert().NoError(err)
		// When
		_, err = p.handler.Buy(ctx, buy)
		// Then
		p.Assert().Error(err)
	})

	p.Run("Given related wager not found then should not save", func() {
		// Given
		buy := &v1.BuyRequest{
			WagerId:     1,
			BuyingPrice: 0.25,
		}

		// When
		_, err := p.handler.Buy(ctx, buy)
		// Then
		p.Assert().Error(err)
	})
}
