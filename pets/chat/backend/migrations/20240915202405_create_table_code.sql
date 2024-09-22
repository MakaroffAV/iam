-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.code (
    id serial not null primary key,
    hash varchar(100) not null,
    code varchar(100) not null,
    user_id int not null references ai_app.user(id),
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ai_app.code;
-- +goose StatementEnd
