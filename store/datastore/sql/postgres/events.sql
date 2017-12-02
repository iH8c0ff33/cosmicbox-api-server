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

-- name: resample-events-timeframe

SELECT
  COUNT(event_timestamp) AS count,
  to_timestamp(floor(Extract(EPOCH FROM event_timestamp) / extract(EPOCH FROM $1::INTERVAL)) *
               extract(EPOCH FROM $1::INTERVAL)) AS intvl
FROM events
WHERE event_timestamp > $2 AND event_timestamp < $3
GROUP BY intvl
ORDER BY intvl ASC;