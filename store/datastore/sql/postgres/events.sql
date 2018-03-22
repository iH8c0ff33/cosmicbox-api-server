-- name: event-find-timestamp

SELECT
  event_id,
  event_timestamp,
  event_pressure
FROM events
WHERE event_timestamp = $1
ORDER BY event_timestamp ASC;

-- name: event-find-range

SELECT
  event_id,
  event_timestamp,
  event_pressure
FROM events
WHERE event_timestamp >= $1
      AND event_timestamp < $2
ORDER BY event_timestamp ASC;

-- name: event-delete

DELETE FROM events
WHERE event_id = $1;

-- name: resample-events-timeframe

SELECT
  to_timestamp(
    floor(
      extract(EPOCH FROM event_timestamp) /
      extract(EPOCH FROM $1::INTERVAL)
    ) *
    extract(EPOCH FROM $1::INTERVAL)
  )                       AS start_time,
  COUNT(event_timestamp)  AS count,
  AVG(event_pressure)     AS avg_press  
FROM events
WHERE event_timestamp > $2 AND event_timestamp < $3
GROUP BY start_time
ORDER BY start_time ASC;