package handler

import (
	"context"
	"encoding/json"
	"sync"

	pbExample "github.com/dotdak/exchange-system/proto"
	v1 "github.com/dotdak/exchange-system/proto/v1"
	"github.com/dotdak/exchange-system/repo"
	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

// Handler implements the protobuf interface
type Handler struct {
	wagerRepo repo.WagerRepo
	buyRepo   repo.BuyRepo
}

// CreateWager implements v1.WagerServiceServer.
func (h *Handler) CreateWager(context.Context, *v1.CreateWagerRequest) (*v1.CreateWagerResponse, error) {
	panic("unimplemented")
}

// ListWagers implements v1.WagerServiceServer.
func (h *Handler) ListWagers(ctx context.Context, req *v1.ListWagersRequest) (*structpb.ListValue, error) {
	buf, err := json.Marshal(req)
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
func (h *Handler) Buy(context.Context, *v1.BuyRequest) (*v1.BuyResponse, error) {
	panic("unimplemented")
}

// New initializes a new Handler struct.
func NewRepo(
	wagerRepo repo.WagerRepo,
	buyRepo repo.BuyRepo,
) *Handler {
	return &Handler{
		wagerRepo: wagerRepo,
		buyRepo:   buyRepo,
	}
}

// Backend implements the protobuf interface
type Backend struct {
	mu    *sync.RWMutex
	users []*pbExample.User
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

// AddUser adds a user to the in-memory store.
func (b *Backend) AddUser(ctx context.Context, _ *pbExample.AddUserRequest) (*pbExample.User, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	user := &pbExample.User{
		Id: uuid.Must(uuid.NewV4()).String(),
	}
	b.users = append(b.users, user)

	return user, nil
}

// ListUsers lists all users in the store.
func (b *Backend) ListUsers(context.Context, *pbExample.ListUsersRequest) (*structpb.ListValue, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	i := make([]interface{}, 0, len(b.users))
	for _, v := range b.users {
		i = append(i, v)
	}

	buf, err := json.Marshal(b.users)
	if err != nil {
		return nil, err
	}

	v := structpb.ListValue{}
	if err := v.UnmarshalJSON(buf); err != nil {
		return nil, err
	}

	return &v, nil
}
