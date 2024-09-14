-- +goose Up
-- +goose StatementBegin
create schema if not exists housing;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema if exists housing;
-- +goose StatementEnd
