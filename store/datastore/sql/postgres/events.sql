-- name: event-find-timestamp

SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp = $1
ORDER BY event_timestamp ASC;

-- name: event-find-range

SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp >= $1
      AND event_timestamp < $2
ORDER BY event_timestamp ASC;

-- name: event-delete

DELETE FROM events
WHERE event_id = $1;
