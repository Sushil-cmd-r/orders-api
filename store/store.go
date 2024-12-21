package store

import (
	"context"

	"github.com/sushil-cmd-r/orders-api/db"
	"github.com/sushil-cmd-r/orders-api/store/model"
)

type Store struct {
	Orders interface {
		Select(ctx context.Context) ([]model.Order, error)
		SelectById(ctx context.Context, id int64) (*model.Order, error)
		Insert(ctx context.Context, order *model.Order) error
		UpdateById(ctx context.Context, id int64, order *model.Order) error
		DeleteById(ctx context.Context, id int64) error
	}
}

func Init(db *db.DB) *Store {
	return &Store{
		Orders: &orderStore{conn: db.Conn()},
	}
}
