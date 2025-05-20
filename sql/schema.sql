CREATE TABLE IF NOT EXISTS events (
  -- PRIMARY KEY (source/user, ID)
  id   INTEGER PRIMARY KEY,
  name TEXT    NOT NULL,
  description TEXT NOT NULL,
  recurrence JSON,
  -- start_date and end_date for all-day events
  -- NOTE: This could be a boolean flag that turns that on/off?
  --       Or two seperate fields?
  start_time INTEGER NOT NULL,
  end_time INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS users (
  user_id   TEXT PRIMARY KEY,
  username TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_active TIMESTAMP
);

CREATE TABLE command_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT,
    guild_id TEXT,
    command TEXT NOT NULL,
    executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
