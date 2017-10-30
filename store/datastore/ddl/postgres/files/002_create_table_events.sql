-- name: create-table-events

CREATE TABLE IF NOT EXISTS events (
  event_id        SERIAL PRIMARY KEY,
  event_timestamp TIMESTAMP
)