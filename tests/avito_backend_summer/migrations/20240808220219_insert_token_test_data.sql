-- +goose Up
-- +goose StatementBegin
insert into housing.token (
	value,
	user_id,
	created
) values (
	'test_token',
	1,
	'2024-08-08 22:00:00'
);
insert into housing.token (
	value,
	user_id,
	created
) values (
	'test_token_moderator',
	2,
	'2024-08-08 22:00:00'
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from housing.token where value = 'test_token';
delete from housing.token where value = 'test_token_moderator';
-- +goose StatementEnd
