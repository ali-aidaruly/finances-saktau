BEGIN;

CREATE TABLE category (
    id BIGINT CONSTRAINT category_pk PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_telegram_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    category_origin_typed TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    deleted_at TIMESTAMP
);

CREATE TRIGGER on_update
    BEFORE UPDATE
    ON category
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

CREATE UNIQUE INDEX ON category (user_telegram_id, name);

COMMIT;