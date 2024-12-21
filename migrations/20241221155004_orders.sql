-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  customer_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW (),
  shipped_at TIMESTAMP,
  completed_at TIMESTAMP
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;

-- +goose StatementEnd
