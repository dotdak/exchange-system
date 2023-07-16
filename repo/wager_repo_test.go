package repo

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/dotdak/exchange-system/dao"
	"gorm.io/gorm"
)

func TestNewWagerRepo(t *testing.T) {
	type args struct {
		db     *gorm.DB
		logger *log.Logger
	}
	tests := []struct {
		name string
		args args
		want WagerRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWagerRepo(tt.args.db, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWagerRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWagerRepoImpl_Create(t *testing.T) {
	type fields struct {
		db     *gorm.DB
		logger *log.Logger
	}
	type args struct {
		ctx   context.Context
		wager *dao.Wager
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dao.Wager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WagerRepoImpl{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.Create(tt.args.ctx, tt.args.wager)
			if (err != nil) != tt.wantErr {
				t.Errorf("WagerRepoImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WagerRepoImpl.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWagerRepoImpl_Update(t *testing.T) {
	type fields struct {
		db     *gorm.DB
		logger *log.Logger
	}
	type args struct {
		ctx   context.Context
		wager *dao.Wager
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dao.Wager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WagerRepoImpl{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.Update(tt.args.ctx, tt.args.wager)
			if (err != nil) != tt.wantErr {
				t.Errorf("WagerRepoImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WagerRepoImpl.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWagerRepoImpl_List(t *testing.T) {
	type fields struct {
		db     *gorm.DB
		logger *log.Logger
	}
	type args struct {
		ctx    context.Context
		offset uint
		limit  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dao.Wager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WagerRepoImpl{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.List(tt.args.ctx, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("WagerRepoImpl.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WagerRepoImpl.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWagerRepoImpl_Get(t *testing.T) {
	type fields struct {
		db     *gorm.DB
		logger *log.Logger
	}
	type args struct {
		ctx context.Context
		id  uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dao.Wager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WagerRepoImpl{
				db:     tt.fields.db,
				logger: tt.fields.logger,
			}
			got, err := r.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WagerRepoImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WagerRepoImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
