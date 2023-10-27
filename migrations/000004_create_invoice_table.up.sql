BEGIN;

CREATE TABLE invoice (
    id BIGINT CONSTRAINT invoice_pk PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_telegram_id BIGINT NOT NULL,
    category_id INTEGER NOT NULL,
    amount DECIMAL NOT NULL,
    currency currency_type NOT NULL,
    description TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    deleted_at TIMESTAMP
);

CREATE TRIGGER on_update
    BEFORE UPDATE
    ON invoice
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

COMMIT;