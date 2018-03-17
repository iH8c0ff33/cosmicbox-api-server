-- name: alter-table-events-add-pressure

ALTER TABLE events
  ADD COLUMN event_pressure FLOAT;