package mysql

// Lookup returns the named statement.
func Lookup(name string) string {
	return index[name]
}

var index = map[string]string{
	"count-users":          countUsers,
	"event-find-timestamp": eventFindTimestamp,
	"event-find-range":     eventFindRange,
	"event-delete":         eventDelete,
	"user-find":            userFind,
	"user-find-login":      userFindLogin,
	"user-update":          userUpdate,
	"user-delete":          userDelete,
}

var countUsers = `
SELECT count(1)
FROM users
`

var eventFindTimestamp = `
SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp = ?
ORDER BY event_timestamp ASC;
`

var eventFindRange = `
SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp >= ?
      AND event_timestamp < ?
ORDER BY event_timestamp ASC;
`

var eventDelete = `
DELETE FROM events
WHERE event_id = ?
`

var userFind = `
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
`

var userFindLogin = `
SELECT
  user_id,
  user_login,
  user_token,
  user_secret,
  user_expiry,
  user_email,
  user_hash
FROM users
WHERE user_login = ?
LIMIT 1;
`

var userUpdate = `
UPDATE users
SET
  user_token  = ?,
  user_secret = ?,
  user_expiry = ?,
  user_email  = ?,
  user_hash   = ?
WHERE user_id = ?;
`

var userDelete = `
DELETE FROM users
WHERE user_id = ?;
`
