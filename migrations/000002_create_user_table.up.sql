BEGIN;

CREATE TYPE currency_type AS ENUM ('KZT');

CREATE TABLE user (
    telegram_id BIGINT CONSTRAINT users_pk PRIMARY KEY,
    telegram_username TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT,
    default_currency currency_type NOT NULL DEFAULT 'KZT',

    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX ON user (telegram_username);

CREATE TRIGGER on_update
    BEFORE UPDATE
    ON user
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

COMMIT;