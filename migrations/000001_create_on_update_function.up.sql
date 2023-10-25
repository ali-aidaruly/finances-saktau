BEGIN;

CREATE OR REPLACE FUNCTION set_updated_at() RETURNS TRIGGER
LANGUAGE plpgsql
AS
$$
BEGIN
    IF (NEW != OLD) THEN
        NEW.updated_at = CURRENT_TIMESTAMP AT TIME ZONE 'utc';
RETURN NEW;
END IF;
RETURN OLD;
END;
$$;

COMMIT;
