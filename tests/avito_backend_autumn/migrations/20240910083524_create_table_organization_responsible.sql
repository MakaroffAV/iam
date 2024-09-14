-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists organization_responsible (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization(id) ON DELETE CASCADE,
    user_id UUID REFERENCES employee(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists organization_responsible;
-- +goose StatementEnd
