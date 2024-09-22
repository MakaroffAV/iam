-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.token (
    id serial not null primary key,
    value varchar(100) not null,
    code_id int not null references ai_app.code(id),
    is_active bool not null default true,
    expires timestamp not null default now() + interval '24 hour',
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ai_app.token;
-- +goose StatementEnd
