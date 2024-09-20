-- +goose Up
CREATE TABLE
    items (
        id bigserial NOT NULL,
        title text NOT NULL,
        amount int NOT NULL,
        quantity int NOT NULL,
        status text,
        PRIMARY KEY (id)
    );

-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS items;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd