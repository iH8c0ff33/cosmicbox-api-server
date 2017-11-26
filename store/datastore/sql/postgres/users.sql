-- name: user-find

SELECT
  user_id,
  user_login,
  user_token,
  user_secret,
  user_expiry,
  user_email,
  user_hash
FROM users
ORDER BY user_login ASC;

-- name: user-find-login

SELECT
  user_id,
  user_login,
  user_token,
  user_secret,
  user_expiry,
  user_email,
  user_hash
FROM users
WHERE user_login = $1
LIMIT 1;

-- name: user-update

UPDATE users
SET
  user_token  = $1,
  user_secret = $2,
  user_expiry = $3,
  user_email  = $4,
  user_hash   = $5
WHERE user_id = $6;

-- name: user-delete

DELETE FROM users
WHERE user_id = $1;
