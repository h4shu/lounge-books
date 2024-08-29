CREATE TABLE IF NOT EXISTS book_tagging (
  id serial NOT NULL PRIMARY KEY,
  book_id integer references books(id),
  tag_id integer references tags(id),
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);