-- +goose Up
-- +goose StatementBegin
create table if not exists tender_archive (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tender_id uuid references tender(id),
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
SELECT 'down SQL query';
-- +goose StatementEnd
