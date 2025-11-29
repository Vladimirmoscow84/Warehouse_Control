BEGIN;

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL UNIQUE CHECK(role_name IN('admin', 'manager', 'viewer')),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100),
    role_id INT NOT NULL REFERENCES roles(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(50) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    price DECIMAL(10,2) NOT NULL DEFAULT 0 CHECK (price >= 0),
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS item_history (
    id SERIAL PRIMARY KEY,
    item_id INT NOT NULL REFERENCES items(id),
    action_type VARCHAR(10) NOT NULL CHECK (action_type IN('insert', 'update', 'delete')),
    old_value JSONB,
    new_value JSONB,
    changed_by INT REFERENCES users(id), 
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION fn_item_before_update() RETURNS TRIGGER AS $$
BEGIN
  IF TG_OP = 'UPDATE' THEN
    NEW.version := COALESCE(OLD.version, 0) + 1;
    NEW.updated_at := CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_item_before_update ON items;
CREATE TRIGGER trg_item_before_update
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE PROCEDURE fn_item_before_update();


CREATE OR REPLACE FUNCTION fn_item_audit() RETURNS TRIGGER AS $$
DECLARE
  v_user_text TEXT;
  v_user_id  INT := NULL;
BEGIN
  BEGIN
    v_user_text := current_setting('app.current_user', true);
  EXCEPTION WHEN OTHERS THEN
    v_user_text := NULL;
  END;

  IF v_user_text IS NOT NULL THEN
    BEGIN
      v_user_id := v_user_text::INT;
    EXCEPTION WHEN OTHERS THEN
      v_user_id := NULL;
    END;
  END IF;

  IF (TG_OP = 'INSERT') THEN
    INSERT INTO item_history(item_id, action_type, new_value, changed_by)
    VALUES (NEW.id, 'insert', to_jsonb(NEW.*), v_user_id);
    RETURN NEW;

  ELSIF (TG_OP = 'UPDATE') THEN
    INSERT INTO item_history(item_id, action_type, old_value, new_value, changed_by)
    VALUES (OLD.id, 'update', to_jsonb(OLD.*), to_jsonb(NEW.*), v_user_id);
    RETURN NEW;

  ELSIF (TG_OP = 'DELETE') THEN
    INSERT INTO item_history(item_id, action_type, old_value, changed_by)
    VALUES (OLD.id, 'delete', to_jsonb(OLD.*), v_user_id);
    RETURN OLD;
  END IF;

  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_item_audit ON items;
CREATE TRIGGER trg_item_audit
AFTER INSERT OR UPDATE OR DELETE ON items
FOR EACH ROW
EXECUTE PROCEDURE fn_item_audit();

COMMIT;
