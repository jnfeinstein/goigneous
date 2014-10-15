-- +goose Up
CREATE TABLE documents
(
  id SERIAL PRIMARY KEY,
  content JSON NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS documents;
