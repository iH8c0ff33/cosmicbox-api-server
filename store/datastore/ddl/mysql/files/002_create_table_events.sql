-- name: create-table-events

CREATE TABLE IF NOT EXISTS events (
  event_id        INTEGER PRIMARY KEY AUTO_INCREMENT,
  event_timestamp TIMESTAMP(6)
)