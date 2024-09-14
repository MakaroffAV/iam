-- +goose Up
-- +goose StatementBegin
insert into housing.house (
	year,
	address,
	developer,
	created,
	updated
) values (
	2001,
	'test address',
	'test developer',
	'2024-08-08 15:30',
	'2024-08-08 15:30'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from housing.house where address = 'test address';
-- +goose StatementEnd
