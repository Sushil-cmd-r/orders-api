-- +goose Up
-- +goose StatementBegin
ALTER TABLE line_items ADD CONSTRAINT fk_orders_line_items FOREIGN KEY (order_id) REFERENCES orders (id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE line_items
DROP CONSTRAINT fk_orders_line_items;

-- +goose StatementEnd
