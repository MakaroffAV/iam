-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.user (
    id serial not null primary key,
    email varchar(200) unique,
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ai_app.user;
-- +goose StatementEnd
