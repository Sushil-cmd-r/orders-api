package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sushil-cmd-r/orders-api/store/model"
)

var ErrNotExist = errors.New("resource not found")

type orderStore struct {
	conn *sql.DB
}

func (s *orderStore) Select(ctx context.Context) ([]model.Order, error) {
	return nil, nil
}

func (s *orderStore) SelectById(ctx context.Context, id int64) (*model.Order, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("select error: %w", err)
	}

	defer tx.Rollback()
	order := model.Order{}
	q1 := `SELECT id, customer_id, created_at, shipped_at, completed_at FROM orders WHERE id = $1;`

	err = tx.QueryRowContext(ctx, q1, id).Scan(&order.Id, &order.CustomerId, &order.CreatedAt, &order.ShippedAt, &order.CompletedAt)

	if errors.Is(sql.ErrNoRows, err) {
		return nil, ErrNotExist
	} else if err != nil {
		return nil, fmt.Errorf("select error: %w", err)
	}

	q2 := `SELECT item_id, price, quantity FROM line_items WHERE order_id = $1;`
	rows, err := tx.QueryContext(ctx, q2, id)
	if err != nil {
		return nil, fmt.Errorf("select error: %w", err)
	}

	var lineItems []model.LineItem
	for rows.Next() {
		var lineItem model.LineItem
		if err := rows.Scan(&lineItem.ItemId, &lineItem.Price, &lineItem.Quantity); err != nil {
			return nil, fmt.Errorf("select error: %w", err)
		}
		lineItems = append(lineItems, lineItem)
	}

	order.LineItems = lineItems
	tx.Commit()
	return &order, nil
}

func (s *orderStore) Insert(ctx context.Context, order *model.Order) error {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("insert error: %w", err)
	}

	defer tx.Rollback()
	q1 := `INSERT INTO orders (customer_id) VALUES ($1) RETURNING id;`

	var orderId int
	err = tx.QueryRowContext(ctx, q1, order.CustomerId).Scan(&orderId)
	if err != nil {
		return fmt.Errorf("insert error: %w", err)
	}

	q2 := `INSERT INTO line_items (price, quantity, order_id) VALUES `
	vals := make([]any, 0)
	for idx, item := range order.LineItems {
		i := idx*3 + 1
		if idx == len(order.LineItems)-1 {
			q2 += fmt.Sprintf("($%d, $%d, $%d);", i, i+1, i+2)
		} else {
			q2 += fmt.Sprintf("($%d, $%d, $%d),", i, i+1, i+2)
		}

		vals = append(vals, item.Price, item.Quantity, orderId)
	}

	_, err = tx.Exec(q2, vals...)
	if err != nil {
		return fmt.Errorf("insert error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("insert error: %w", err)
	}

	return nil
}

func (s *orderStore) UpdateById(ctx context.Context, id int64, order *model.Order) error {
	return nil
}

func (s *orderStore) DeleteById(ctx context.Context, id int64) error {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	defer tx.Rollback()

	q := `DELETE FROM orders WHERE id = $1;`
	_, err = tx.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}
	tx.Commit()

	return nil
}
