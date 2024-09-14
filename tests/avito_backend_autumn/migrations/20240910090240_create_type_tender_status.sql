-- +goose Up
-- +goose StatementBegin
create type tender_status as enum (
    'Created',
    'Published',
    'Closed'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists tender_status;
-- +goose StatementEnd
