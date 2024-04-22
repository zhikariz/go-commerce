BEGIN;

ALTER TABLE users
DROP COLUMN created_at,
DROP COLUMN updated_at,
DROP COLUMN deleted_at;

COMMIT;