-- +goose Up
-- +goose StatementBegin
create table if not exists housing.flat (
	id bigint not null,
	price int not null,
	rooms int not null,
	house_id int not null references housing.house(id),
	status varchar(200) not null default 'created',
	unique (id, house_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists housing.flat;
-- +goose StatementEnd
