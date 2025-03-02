package db

import (
  "context"
  "database/sql"
  "fmt"
  rawsql "github.com/ohhfishal/schedule/sql"
  _ "modernc.org/sqlite"
)

func Connect(ctx context.Context, driver, source string) (*Queries, error) {
  if driver != `sqlite` {
    return nil, fmt.Errorf(`driver not supported: %s`, driver)

  }
  conn, err := sql.Open(driver, source)
  if err != nil {
    return nil, fmt.Errorf(`connecting to database: %w`, err)
  }

  if _, err := conn.ExecContext(ctx, rawsql.Schema); err != nil {
    return nil, fmt.Errorf(`running schema: %w`, err)
  }
  return New(conn), nil
}
