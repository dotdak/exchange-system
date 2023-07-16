package handler

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"time"

	"github.com/dotdak/exchange-system/dao"
	"github.com/dotdak/exchange-system/pkg/es_errors"
	"github.com/dotdak/exchange-system/pkg/utils"
	v1 "github.com/dotdak/exchange-system/proto/v1"
	"github.com/dotdak/exchange-system/repo"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const DefaultLimit = 10

type Handler interface {
	v1.WagerServiceServer
	v1.BuyServiceServer
}

// Handler implements the protobuf interface
type HandlerImpl struct {
	v1.UnimplementedBuyServiceServer
	v1.UnimplementedWagerServiceServer

	logger    *log.Logger
	wagerRepo repo.WagerRepo
	buyRepo   repo.BuyRepo
}

// New initializes a new Handler struct.
func NewHandler(
	wagerRepo repo.WagerRepo,
	buyRepo repo.BuyRepo,
) Handler {
	return &HandlerImpl{
		logger:    log.Default(),
		wagerRepo: wagerRepo,
		buyRepo:   buyRepo,
	}
}

// CreateWager implements v1.WagerServiceServer.
func (h *HandlerImpl) CreateWager(ctx context.Context, req *v1.CreateWagerRequest) (*v1.CreateWagerResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	wager := dao.Wager{
		TotalWagerValue:     req.TotalWagerValue,
		Odds:                req.Odds,
		SellingPercentage:   req.SellingPercentage,
		SellingPrice:        req.SellingPrice,
		CurrentSellingPrice: req.SellingPrice,
		PlacedAt:            time.Now().UTC(),
	}

	res, err := h.wagerRepo.Create(ctx, &wager)
	if err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

// ListWagers implements v1.WagerServiceServer.
func (h *HandlerImpl) ListWagers(ctx context.Context, req *v1.ListWagersRequest) (*structpb.ListValue, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	limit := utils.Any(req.Limit, DefaultLimit)
	page := utils.Any(req.Page, 1)
	offset := (page - 1) * limit

	wagers, err := h.wagerRepo.List(ctx, uint(offset), uint(limit))
	if err != nil {
		return nil, err
	}

	buf, err := json.Marshal(wagers)
	if err != nil {
		return nil, err
	}

	v := structpb.ListValue{}
	if err := v.UnmarshalJSON(buf); err != nil {
		return nil, err
	}

	return &v, nil
}

// Buy implements v1.BuyServiceServer.
func (h *HandlerImpl) Buy(ctx context.Context, req *v1.BuyRequest) (*v1.BuyResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	wager, err := h.wagerRepo.Get(ctx, req.WagerId)
	if err != nil {
		return nil, err
	}

	if wager.CurrentSellingPrice < req.BuyingPrice {
		h.logger.Printf(
			"CurrentSellingPrice %f < BuyingPrice %f", wager.CurrentSellingPrice, req.BuyingPrice,
		)
		return nil, es_errors.ErrBuyHigherThanSell
	}

	wager.CurrentSellingPrice -= req.BuyingPrice
	wager.AmountSold += req.BuyingPrice
	wager.PercentageSold = uint32(math.Round(wager.AmountSold / wager.SellingPrice * 100))

	wager, err = h.wagerRepo.Update(ctx, wager)
	if err != nil {
		h.logger.Fatalf("update wager failed: %v", err)
		return nil, err
	}

	buy, err := h.buyRepo.Create(ctx, &dao.Buy{
		WagerID:     uint32(wager.ID),
		BuyingPrice: req.BuyingPrice,
		BoughtAt:    time.Now().UTC(),
	})

	if err != nil {
		return nil, err
	}

	return &v1.BuyResponse{
		Id:          uint32(buy.ID),
		WagerId:     buy.WagerID,
		BuyingPrice: buy.BuyingPrice,
		BoughtAt:    timestamppb.New(buy.BoughtAt),
	}, nil
}
