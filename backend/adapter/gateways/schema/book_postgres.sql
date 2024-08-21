CREATE TABLE IF NOT EXISTS books (
  id serial NOT NULL PRIMARY KEY,
  isbn text,
  title text,
  description text,
  cover_link text,
  published_at date,
  author text,
  publisher text,
  page_count integer,
  deleted_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);