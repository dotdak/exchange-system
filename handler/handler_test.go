package handler

import (
	"context"
	"log"
	"reflect"
	"testing"

	v1 "github.com/dotdak/exchange-system/proto/v1"
	"github.com/dotdak/exchange-system/repo"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		wagerRepo repo.WagerRepo
		buyRepo   repo.BuyRepo
		logger    *log.Logger
		db        *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.wagerRepo, tt.args.buyRepo, tt.args.logger, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerImpl_CreateWager(t *testing.T) {
	type fields struct {
		UnimplementedBuyServiceServer   v1.UnimplementedBuyServiceServer
		UnimplementedWagerServiceServer v1.UnimplementedWagerServiceServer
		logger                          *log.Logger
		wagerRepo                       repo.WagerRepo
		buyRepo                         repo.BuyRepo
		db                              *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *v1.CreateWagerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.CreateWagerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				UnimplementedBuyServiceServer:   tt.fields.UnimplementedBuyServiceServer,
				UnimplementedWagerServiceServer: tt.fields.UnimplementedWagerServiceServer,
				logger:                          tt.fields.logger,
				wagerRepo:                       tt.fields.wagerRepo,
				buyRepo:                         tt.fields.buyRepo,
				db:                              tt.fields.db,
			}
			got, err := h.CreateWager(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.CreateWager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerImpl.CreateWager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerImpl_ListWagers(t *testing.T) {
	type fields struct {
		UnimplementedBuyServiceServer   v1.UnimplementedBuyServiceServer
		UnimplementedWagerServiceServer v1.UnimplementedWagerServiceServer
		logger                          *log.Logger
		wagerRepo                       repo.WagerRepo
		buyRepo                         repo.BuyRepo
		db                              *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *v1.ListWagersRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *structpb.ListValue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				UnimplementedBuyServiceServer:   tt.fields.UnimplementedBuyServiceServer,
				UnimplementedWagerServiceServer: tt.fields.UnimplementedWagerServiceServer,
				logger:                          tt.fields.logger,
				wagerRepo:                       tt.fields.wagerRepo,
				buyRepo:                         tt.fields.buyRepo,
				db:                              tt.fields.db,
			}
			got, err := h.ListWagers(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.ListWagers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerImpl.ListWagers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerImpl_Buy(t *testing.T) {
	type fields struct {
		UnimplementedBuyServiceServer   v1.UnimplementedBuyServiceServer
		UnimplementedWagerServiceServer v1.UnimplementedWagerServiceServer
		logger                          *log.Logger
		wagerRepo                       repo.WagerRepo
		buyRepo                         repo.BuyRepo
		db                              *gorm.DB
	}
	type args struct {
		ctx context.Context
		req *v1.BuyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.BuyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HandlerImpl{
				UnimplementedBuyServiceServer:   tt.fields.UnimplementedBuyServiceServer,
				UnimplementedWagerServiceServer: tt.fields.UnimplementedWagerServiceServer,
				logger:                          tt.fields.logger,
				wagerRepo:                       tt.fields.wagerRepo,
				buyRepo:                         tt.fields.buyRepo,
				db:                              tt.fields.db,
			}
			got, err := h.Buy(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandlerImpl.Buy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerImpl.Buy() = %v, want %v", got, tt.want)
			}
		})
	}
}
