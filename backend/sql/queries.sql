-- name: insert-task
INSERT INTO
  tasks (
    task_name,
    start_time,
    end_time,
    work_length,
    break_length
  )
VALUES
  (?, ?, ?, ?, ?)
