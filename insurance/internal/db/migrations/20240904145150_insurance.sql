-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS insurances
(
    status      int,
    active_till timestamp,
    id          text unique,
    price       int
);

CREATE INDEX insurance_id_idx ON insurances USING HASH (id);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
