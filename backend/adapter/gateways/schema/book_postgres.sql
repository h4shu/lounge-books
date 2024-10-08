CREATE TABLE IF NOT EXISTS books (
  id serial NOT NULL PRIMARY KEY,
  isbn text,
  title text NOT NULL,
  description text NOT NULL,
  cover_link text NOT NULL,
  published_year integer NOT NULL,
  published_month integer NOT NULL,
  published_day integer NOT NULL,
  author text NOT NULL,
  publisher text NOT NULL,
  page_count integer NOT NULL,
  deleted_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);