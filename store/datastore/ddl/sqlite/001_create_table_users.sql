-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
  user_id     INTEGER PRIMARY KEY AUTOINCREMENT,
  user_login  TEXT,
  user_token  TEXT,
  user_secret TEXT,
  user_expiry DATETIME,
  user_email  TEXT,
  user_hash   TEXT,
  UNIQUE (user_login)
);