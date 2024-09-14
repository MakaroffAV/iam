-- +goose Up
-- +goose StatementBegin
create table if not exists housing.house (
	id serial not null primary key,
	address varchar(500) not null,
	developer varchar(100) not null,
	year int not null,
	created timestamp not null default current_timestamp,
	updated timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists housing.house;
-- +goose StatementEnd
