package store

import "github.com/sushil-cmd-r/orders-api/store/model"

type orderStore struct{}

func (s *orderStore) Select() ([]model.Order, error) {
	return nil, nil
}

func (s *orderStore) SelectById(id int64) (*model.Order, error) {
	return nil, nil
}

func (s *orderStore) Insert(order *model.Order) error {
	return nil
}

func (s *orderStore) UpdateById(id int64, order *model.Order) error {
	return nil
}

func (s *orderStore) DeleteById(id int64) error {
	return nil
}
