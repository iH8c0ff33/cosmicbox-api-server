-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
  user_id     SERIAL PRIMARY KEY,
  user_login  VARCHAR(250),
  user_token  VARCHAR(500),
  user_secret VARCHAR(500),
  user_expiry TIMESTAMP,
  user_email  VARCHAR(500),
  user_hash   VARCHAR(500),
  UNIQUE (user_login)
);