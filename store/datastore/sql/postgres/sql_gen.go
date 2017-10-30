package postgres

// Lookup returns the named statement.
func Lookup(name string) string {
	return index[name]
}

var index = map[string]string{
	"count-users":          countUsers,
	"count-events":         countEvents,
	"event-find-timestamp": eventFindTimestamp,
	"event-find-range":     eventFindRange,
	"event-delete":         eventDelete,
	"user-find":            userFind,
	"user-find-login":      userFindLogin,
	"user-update":          userUpdate,
	"user-delete":          userDelete,
}

var countUsers = `
SELECT reltuples
FROM pg_class
WHERE relname = 'users';
`

var countEvents = `
SELECT reltuples
FROM pg_class
WHERE relname = 'events'
`

var eventFindTimestamp = `
SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp = $1
ORDER BY event_timestamp ASC;
`

var eventFindRange = `
SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp >= $1
      AND event_timestamp < $2
ORDER BY event_timestamp ASC;
`

var eventDelete = `
DELETE FROM events
WHERE event_id = $1
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
WHERE user_login = $1
LIMIT 1;
`

var userUpdate = `
UPDATE users
SET
  user_token  = $1,
  user_secret = $2,
  user_expiry = $3,
  user_email  = $4,
  user_hash   = $5
WHERE user_id = $6;
`

var userDelete = `
DELETE FROM users
WHERE user_id = $1
`
