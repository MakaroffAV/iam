-- +goose Up
-- +goose StatementBegin
create table if not exists housing.token (
	id serial not null primary key,
	value varchar(100) not null,
	user_id int not null references housing.user (id),
	created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists housing.token;
-- +goose StatementEnd
