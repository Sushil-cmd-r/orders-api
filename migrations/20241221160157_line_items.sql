-- +goose Up
-- +goose StatementBegin
CREATE TABLE line_items (
  item_id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID (),
  price INTEGER NOT NULL,
  quantity INTEGER NOT NULL,
  order_id INTEGER NOT NULL
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE line_items;

-- +goose StatementEnd
