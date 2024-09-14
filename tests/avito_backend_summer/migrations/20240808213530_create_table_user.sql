-- +goose Up
-- +goose StatementBegin
create table if not exists housing.user (
	id serial not null primary key,
	type varchar(100) not null,
	uuid varchar(100) not null,
	email varchar(100) not null,
	password varchar(100) not null,
	dummy bool not null,
	created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists housing.user;
-- +goose StatementEnd
