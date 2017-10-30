-- name: create-table-events

CREATE TABLE IF NOT EXISTS events (
  event_id        INTEGER PRIMARY KEY AUTOINCREMENT,
  event_timestamp TIMESTAMP DATETIME
)