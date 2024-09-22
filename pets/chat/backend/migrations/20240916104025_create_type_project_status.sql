-- +goose Up
-- +goose StatementBegin
create type project_status as enum (
    'Created',
    'Processing',
    'Ready',
    'Deleted'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists project_status;
-- +goose StatementEnd
