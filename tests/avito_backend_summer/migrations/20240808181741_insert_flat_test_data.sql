-- +goose Up
-- +goose StatementBegin
insert into housing.flat (
	id,
	price,
	rooms,
	house_id
) values (
	1,
	192837465,
	2,
	1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from housing.flat where price = 192837465;
-- +goose StatementEnd
