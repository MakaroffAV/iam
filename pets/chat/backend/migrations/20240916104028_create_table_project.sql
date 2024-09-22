-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.project (
    id serial not null primary key,
    uuid varchar(100) not null unique,
    name varchar(100) not null,
    description varchar(500) not null,
    status project_status not null default 'Created',
    user_id int not null references ai_app.user(id),
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if not exists ai_app.project
-- +goose StatementEnd
