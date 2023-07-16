//go:build integration

package integration

import (
	"context"
	"math"
	"time"

	"github.com/dotdak/exchange-system/dao"
	"github.com/dotdak/exchange-system/pkg/utils"
	v1 "github.com/dotdak/exchange-system/proto/v1"
)

func (p *MySqlRepositoryTestSuite) TestBuyFlow() {
	ctx := context.Background()
	p.Run("Given valid buy then should save", func() {
		// Given
		amoundSold := float64(100)
		wager := &dao.Wager{
			BaseModel:           dao.BaseModel{ID: 1},
			TotalWagerValue:     1000,
			Odds:                10,
			SellingPercentage:   10,
			SellingPrice:        300,
			CurrentSellingPrice: 150,
			PercentageSold:      utils.NewPointer[uint32](10),
			AmountSold:          &amoundSold,
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

		currentWager := &dao.Wager{}
		p.Assert().NoError(p.db.First(currentWager, buyRes.WagerId).Error)
		p.Assert().Equal(float64(50), currentWager.CurrentSellingPrice)
		p.Assert().Equal(uint32(math.Round((amoundSold+buy.BuyingPrice)/float64(1000)*100)), *currentWager.PercentageSold)
		p.Assert().Equal(amoundSold+buy.BuyingPrice, *currentWager.AmountSold)
	})

	p.Run("Given buy price greater than current selling price then should not save", func() {
		// Given
		wager := &dao.Wager{
			BaseModel:           dao.BaseModel{ID: 1},
			TotalWagerValue:     100,
			Odds:                1,
			SellingPercentage:   10,
			SellingPrice:        30,
			CurrentSellingPrice: 10,
			PlacedAt:            time.Now(),
		}
		buy := &v1.BuyRequest{
			WagerId:     1,
			BuyingPrice: 10.1,
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

	p.Run("Given buy with no wager id then should not save", func() {
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
