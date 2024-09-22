-- +goose Up
-- +goose StatementBegin
create schema if not exists ai_app;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema if exists ai_app;
-- +goose StatementEnd
