-- +goose Up
-- +goose StatementBegin
create type project_query_status as enum (
    'Created',
    'Processing',
    'Ready',
    'Failed'
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists project_query_status;
-- +goose StatementEnd
