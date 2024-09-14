-- +goose Up
-- +goose StatementBegin
insert into housing.user (
	type,
	uuid,
	email,
	password,
	dummy,
	created
) values (
	'client',
	'test_uuid',
	'test_email',
	'test_password',
	false,
	'2024-08-08 21:30'
);
insert into housing.user (
	type,
	uuid,
	email,
	password,
	dummy,
	created
) values (
	'moderator',
	'test_uuid_moderator',
	'test_password_moderator',
	'test_email_moderator',
	false,
	'2024-08-08 21:30'
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from housing.user where uuid = 'test_uuid';
delete from housing.user where uuid = 'test_uuid_moderator';
-- +goose StatementEnd
