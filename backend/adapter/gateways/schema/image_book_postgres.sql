CREATE TABLE IF NOT EXISTS images_books (
  id serial NOT NULL PRIMARY KEY,
  image_id text NOT NULL,
  book_id integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);