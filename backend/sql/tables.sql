-- name: create-tasks-table
CREATE TABLE IF NOT EXISTS tasks (
  task_id INTEGER PRIMARY KEY AUTOINCREMENT,
  task_name TEXT,
  start_time DATETIME,
  end_time DATETIME,
  work_length INTEGER,
  break_length INTEGER
)
