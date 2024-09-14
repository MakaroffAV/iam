-- +goose Up
-- +goose StatementBegin
CREATE TYPE tender_service AS ENUM (
    'Construction',
    'Delivery',
    'Manufacture'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type if exists tender_service;
-- +goose StatementEnd
