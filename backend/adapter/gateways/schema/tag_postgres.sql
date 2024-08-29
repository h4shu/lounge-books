CREATE TABLE IF NOT EXISTS tags (
  id serial NOT NULL PRIMARY KEY,
  name text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);