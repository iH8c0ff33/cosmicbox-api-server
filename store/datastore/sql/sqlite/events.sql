-- name: event-find-timestamp

SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp = ?
ORDER BY event_timestamp
  ASC;

-- name: event-find-range

SELECT
  event_id,
  event_timestamp
FROM events
WHERE event_timestamp >= ?
      AND event_timestamp < ?
ORDER BY event_timestamp
  ASC;

-- name: event-delete

DELETE FROM events
WHERE event_id = ?;

-- name: resample-events-timeframe

SELECT
  COUNT(event_timestamp)                                                                                  AS count,
  datetime(cast(strftime('%s', event_timestamp) /
                strftime('%s', datetime(datetime(0, 'unixepoch'), replace(?1, 's', ' seconds'))) as int) *
           strftime('%s', datetime(datetime(0, 'unixepoch'), replace(?1, 's', ' seconds'))), 'unixepoch') AS intvl
FROM events
WHERE event_timestamp > ?2 AND event_timestamp < ?3
GROUP BY intvl
ORDER BY intvl
  ASC;