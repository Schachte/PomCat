-- name: insert-task
INSERT INTO
  tasks (
    task_id,
    task_name,
    start_time,
    end_time,
    work_length,
    break_length,
    category,
    ended_early,
    early_termination_time
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?)
