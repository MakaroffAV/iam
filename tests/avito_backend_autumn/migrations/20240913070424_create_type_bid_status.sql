-- +goose Up
-- +goose StatementBegin
create type bid_status as enum (
    'Created',
    'Published',
    'Canceled'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists bid_status;
-- +goose StatementEnd
