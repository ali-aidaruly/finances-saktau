BEGIN;

    DROP TABLE IF EXISTS user;
    DROP TRIGGER IF EXISTS on_update ON user;
    DROP TYPE IF EXISTS currency_type;

COMMIT;