package store

import (
	"context"
	"database/sql"

	"github.com/sushil-cmd-r/orders-api/store/model"
)

type orderStore struct {
	conn *sql.DB
}

func (s *orderStore) Select(ctx context.Context) ([]model.Order, error) {
	return nil, nil
}

func (s *orderStore) SelectById(ctx context.Context, id int64) (*model.Order, error) {
	return nil, nil
}

func (s *orderStore) Insert(ctx context.Context, order *model.Order) error {
	return nil
}

func (s *orderStore) UpdateById(ctx context.Context, id int64, order *model.Order) error {
	return nil
}

func (s *orderStore) DeleteById(ctx context.Context, id int64) error {
	return nil
}
