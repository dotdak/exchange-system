//go:build integration

package integration

import (
	"context"

	"github.com/dotdak/exchange-system/dao"
	v1 "github.com/dotdak/exchange-system/proto/v1"
)

func (p *MySqlRepositoryTestSuite) TestWagerRepo_Create() {
	ctx := context.Background()
	p.Run("Given Invalid SellingPrice then should not save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
		}

		// When
		_, err := p.handler.CreateWager(ctx, newWager)

		//Then
		p.Assert().Error(err)
	})

	p.Run("Given Invalid SellingPercentage then should not save", func() {
		// Given
		newWager := []*v1.CreateWagerRequest{{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 101,
		}, {
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 0,
		}}

		// When
		for _, wager := range newWager {
			_, err := p.handler.CreateWager(ctx, wager)
			//Then
			p.Assert().Error(err)
		}
	})

	p.Run("Given SellingPrice < Total x Percentage then should not save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   2,
			Odds:              1,
			SellingPercentage: 50,
			SellingPrice:      1,
		}

		// When
		_, err := p.handler.CreateWager(ctx, newWager)

		// Then
		p.Assert().Error(err)
	})

	p.Run("Given invalid Odds then should not save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              0,
			SellingPercentage: 1,
			SellingPrice:      1,
		}

		// When
		_, err := p.handler.CreateWager(ctx, newWager)

		// Then
		p.Assert().Error(err)
	})

	p.Run("Given invalid TotalWagerValue", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1,
		}

		// When
		_, err := p.handler.CreateWager(ctx, newWager)

		// Then
		p.Assert().Error(err)
	})

	p.Run("Given invalid SellingPrice monetary format then should not save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1.222,
		}

		// When
		_, err := p.handler.CreateWager(ctx, newWager)

		// Then
		p.Assert().Error(err)
	})

	p.Run("Given invalid CurrentSellingPrice, SellingPrice and AmountSold then should not save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1.22,
		}
		wagerResponse, err := p.handler.CreateWager(ctx, newWager)
		p.Assert().NoError(err)
		wager := &dao.Wager{}
		wager.FromProto(wagerResponse)
		wager.CurrentSellingPrice = 1.21
		// When
		err = p.db.Save(wager).Error
		// Then
		p.Assert().NoError(err)
	})

	p.Run("Given valid wager then should save", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1.22,
		}
		// When
		wager, err := p.handler.CreateWager(ctx, newWager)
		// Then
		p.Assert().NoError(err)

		// When
		actualWager, err := p.wagerRepo.Get(ctx, wager.Id)
		// Then
		p.Assert().NoError(err)
		p.Assert().Equal(uint(1), actualWager.ID)
		p.Assert().Equal(1.22, actualWager.SellingPrice)
		p.Assert().Nil(actualWager.AmountSold)
		p.Assert().Nil(actualWager.PercentageSold)
	})
}

func (p *MySqlRepositoryTestSuite) TestWagerRepo_List() {
	ctx := context.Background()
	p.Run("Success", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1.22,
		}
		// When
		_, err := p.handler.CreateWager(ctx, newWager)
		// Then
		p.Assert().NoError(err)
		// When
		_, err = p.handler.CreateWager(ctx, newWager)
		// Then
		p.Assert().NoError(err)
		// When
		_, err = p.handler.CreateWager(ctx, newWager)
		// Then
		p.Assert().NoError(err)

		// When
		actualWagers, err := p.wagerRepo.List(ctx, 1, 10)
		// Then
		p.Assert().NoError(err)
		p.Assert().Equal(2, len(actualWagers))
	})
}

func (p *MySqlRepositoryTestSuite) TestWagerRepo_Get() {
	ctx := context.Background()

	p.Run("Success", func() {
		// Given
		newWager := &v1.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      1.22,
		}
		// When
		wager, err := p.handler.CreateWager(ctx, newWager)
		// Then
		p.Assert().NoError(err)

		// When
		actualWager, err := p.wagerRepo.Get(ctx, uint32(wager.Id))
		// Then
		p.Assert().NoError(err)
		p.Assert().Equal(uint(1), actualWager.ID)
		p.Assert().Equal(1.22, actualWager.SellingPrice)
	})
}
