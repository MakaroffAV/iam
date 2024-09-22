-- +goose Up
-- +goose StatementBegin
create table if not exists ai_app.project_query (
    id serial not null primary key,
    uuid varchar(100) not null,
    project_id int not null references ai_app.project(id),
    query varchar(500) not null,
    enrichment bool not null,
    status project_query_status not null default 'Created',
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ai_app.project_query;
-- +goose StatementEnd
