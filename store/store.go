package store

import (
	"github.com/sushil-cmd-r/orders-api/store/model"
)

type Store struct {
	Orders interface {
		Select() ([]model.Order, error)
		SelectById(id int64) (*model.Order, error)
		Insert(order *model.Order) error
		UpdateById(id int64, order *model.Order) error
		DeleteById(id int64) error
	}
}

func Init() *Store {
	return &Store{
		Orders: &orderStore{},
	}
}
