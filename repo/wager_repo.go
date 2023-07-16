package repo

import (
	"context"

	v1 "github.com/dotdak/exchange-system/proto/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type WagerRepo interface {
	CreateWager(context.Context, *v1.CreateWagerRequest) (*v1.CreateWagerResponse, error)
	ListWagers(context.Context, *v1.ListWagersRequest) (*structpb.ListValue, error)
}
