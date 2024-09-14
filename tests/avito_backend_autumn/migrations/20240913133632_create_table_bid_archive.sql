-- +goose Up
-- +goose StatementBegin
create table if not exists bid_archive (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id uuid not null references bid(id),
    name varchar(100) not null,
    description text not null,
    status bid_status not null default 'Created',
    version int not null default 1,
    tender_id uuid references tender(id),
    author_type bid_author_type not null,
    author_id uuid references employee(id),
    created timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists bid_archive;
-- +goose StatementEnd
