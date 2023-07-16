package repo

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/dotdak/exchange-system/dao"
	"gorm.io/gorm"
)

func TestNewBuyRepo(t *testing.T) {
	type args struct {
		db     *gorm.DB
		logger *log.Logger
	}
	tests := []struct {
		name string
		args args
		want BuyRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuyRepo(tt.args.db, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuyRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuyRepoImpl_Create(t *testing.T) {
	type fields struct {
		db     *gorm.DB
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		buy *dao.Buy
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dao.Buy
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuyRepoImpl{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.Create(tt.args.ctx, tt.args.buy)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyRepoImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuyRepoImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
