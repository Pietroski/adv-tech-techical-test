-- name: GetTask :one
SELECT *
FROM tasks
WHERE task_id = $1
LIMIT 1;

-- name: ListTasks :many
SELECT *
FROM tasks
ORDER BY time;

-- name: ListPaginatedTasks :many
SELECT *
FROM tasks
ORDER BY table_id
LIMIT $1 OFFSET $2;

-- name: CreateTask :one
INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE user_id = $1;

-- name: GetTasksByUserID :many
SELECT *
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = &1;

-- name: GetPaginatedTasksByUserID :many
SELECT *
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = $1
ORDER BY t.table_id
LIMIT $2 OFFSET $3;

-- name: GetTasksByUserEmail :many
SELECT *
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.email = &1;

-- name: GetPaginatedTasksByUserEmail :many
SELECT *
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = $1
ORDER BY t.table_id
LIMIT $2 OFFSET $3;
