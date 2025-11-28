BEGIN;

DROP TRIGGER IF EXISTS trg_item_audit ON items;
DROP FUNCTION IF EXISTS fn_item_audit();

DROP TRIGGER IF EXISTS trg_item_before_update ON items;
DROP FUNCTION IF EXISTS fn_item_before_update();

DROP TABLE IF EXISTS item_history;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

COMMIT;
