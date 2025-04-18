CREATE TABLE IF NOT EXISTS events (
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  description TEXT NOT NULL,
  recurrence JSON,
  start_time INTEGER NOT NULL,
  end_time INTEGER NOT NULL DEFAULT 0
);
