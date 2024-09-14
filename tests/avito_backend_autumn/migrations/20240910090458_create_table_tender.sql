-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists tender (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description varchar(500) not null,
    service_type tender_service not null,
    status tender_status not null,
    employee_id uuid not null REFERENCES employee(id),
    version int not null default 1,
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tender;
-- +goose StatementEnd
