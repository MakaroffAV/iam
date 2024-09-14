-- +goose Up
-- +goose StatementBegin
alter table tender alter status set default 'Created';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table tender alter status drop default;
-- +goose StatementEnd
