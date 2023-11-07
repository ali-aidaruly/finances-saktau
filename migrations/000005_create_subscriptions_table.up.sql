BEGIN;

CREATE TYPE subscription_payment_type AS ENUM ('monthly', 'annual');

CREATE TABLE subscription (
                         id BIGINT CONSTRAINT subscription_pk PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
                         user_telegram_id BIGINT NOT NULL,
                         name TEXT NOT NULL ,
                         amount DECIMAL NOT NULL,
                         currency currency_type NOT NULL,
                         payment_interval subscription_payment_type NOT NULL,
                         description TEXT,

                         created_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
                         updated_at TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
                         deleted_at TIMESTAMP
);

ALTER TYPE currency_type ADD VALUE 'USD';

CREATE TRIGGER on_update
    BEFORE UPDATE
    ON subscription
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at();

COMMIT;