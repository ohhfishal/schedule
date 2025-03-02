CREATE TABLE IF NOT EXISTS events (
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  description TEXT NOT NULL,
  start_time INTEGER NOT NULL,
  end_time INTEGER DEFAULT 0
);
