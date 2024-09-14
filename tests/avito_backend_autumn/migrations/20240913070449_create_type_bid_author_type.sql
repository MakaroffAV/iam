-- +goose Up
-- +goose StatementBegin
create type bid_author_type as enum (
    'User',
    'Organization'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists bid_author_type;
-- +goose StatementEnd
