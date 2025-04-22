CREATE TABLE IF NOT EXISTS events (
  -- PRIMARY KEY (source/user, ID)
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  description TEXT NOT NULL,
  recurrence JSON,
  -- start_date and end_date for all-day events
  start_time INTEGER NOT NULL,
  end_time INTEGER NOT NULL DEFAULT 0
);
