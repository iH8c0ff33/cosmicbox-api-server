-- name: count-users

SELECT reltuples
FROM pg_class
WHERE relname = 'users';

-- name: count-events

SELECT reltuples
FROM pg_class
WHERE relname = 'events'