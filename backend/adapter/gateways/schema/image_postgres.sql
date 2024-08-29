CREATE TABLE IF NOT EXISTS images_books (
  id serial NOT NULL PRIMARY KEY,
  data_id text NOT NULL,
  s3_key text NOT NULL,
  format text NOT NULL,
  original_url text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);