-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.project_file (
    id serial not null primary key,
    project_id int not null references ai_app.project(id),
    uuid varchar(100) not null unique,
    name varchar(100) not null,
    body bytea not null,
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ai_app.project_file;
-- +goose StatementEnd
