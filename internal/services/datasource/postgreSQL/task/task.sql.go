// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package task

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (task_id, user_id, data_range, reminder_period)
VALUES ($1, $2, $3, $4)
RETURNING table_id, task_id, user_id, data_range, reminder_period, created_at
`

type CreateTaskParams struct {
	TaskID         uuid.UUID     `json:"taskID"`
	UserID         uuid.UUID     `json:"userID"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Tasks, error) {
	row := q.queryRow(ctx, q.createTaskStmt, createTask,
		arg.TaskID,
		arg.UserID,
		arg.DataRange,
		arg.ReminderPeriod,
	)
	var i Tasks
	err := row.Scan(
		&i.TableID,
		&i.TaskID,
		&i.UserID,
		&i.DataRange,
		&i.ReminderPeriod,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE user_id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, userID uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, userID)
	return err
}

const getPaginatedTasksByUserEmail = `-- name: GetPaginatedTasksByUserEmail :many
SELECT u.table_id, u.user_id, email, first_name, last_name, u.created_at, t.table_id, task_id, t.user_id, data_range, reminder_period, t.created_at
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = $1
ORDER BY t.table_id
LIMIT $2 OFFSET $3
`

type GetPaginatedTasksByUserEmailParams struct {
	UserID uuid.UUID `json:"userID"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetPaginatedTasksByUserEmailRow struct {
	TableID        sql.NullInt64 `json:"tableID"`
	UserID         uuid.UUID     `json:"userID"`
	Email          string        `json:"email"`
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	CreatedAt      sql.NullTime  `json:"createdAt"`
	TableID_2      sql.NullInt64 `json:"tableID2"`
	TaskID         uuid.UUID     `json:"taskID"`
	UserID_2       uuid.UUID     `json:"userID2"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
	CreatedAt_2    sql.NullTime  `json:"createdAt2"`
}

func (q *Queries) GetPaginatedTasksByUserEmail(ctx context.Context, arg GetPaginatedTasksByUserEmailParams) ([]GetPaginatedTasksByUserEmailRow, error) {
	rows, err := q.query(ctx, q.getPaginatedTasksByUserEmailStmt, getPaginatedTasksByUserEmail, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPaginatedTasksByUserEmailRow{}
	for rows.Next() {
		var i GetPaginatedTasksByUserEmailRow
		if err := rows.Scan(
			&i.TableID,
			&i.UserID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.TableID_2,
			&i.TaskID,
			&i.UserID_2,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaginatedTasksByUserID = `-- name: GetPaginatedTasksByUserID :many
SELECT u.table_id, u.user_id, email, first_name, last_name, u.created_at, t.table_id, task_id, t.user_id, data_range, reminder_period, t.created_at
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = $1
ORDER BY t.table_id
LIMIT $2 OFFSET $3
`

type GetPaginatedTasksByUserIDParams struct {
	UserID uuid.UUID `json:"userID"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetPaginatedTasksByUserIDRow struct {
	TableID        sql.NullInt64 `json:"tableID"`
	UserID         uuid.UUID     `json:"userID"`
	Email          string        `json:"email"`
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	CreatedAt      sql.NullTime  `json:"createdAt"`
	TableID_2      sql.NullInt64 `json:"tableID2"`
	TaskID         uuid.UUID     `json:"taskID"`
	UserID_2       uuid.UUID     `json:"userID2"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
	CreatedAt_2    sql.NullTime  `json:"createdAt2"`
}

func (q *Queries) GetPaginatedTasksByUserID(ctx context.Context, arg GetPaginatedTasksByUserIDParams) ([]GetPaginatedTasksByUserIDRow, error) {
	rows, err := q.query(ctx, q.getPaginatedTasksByUserIDStmt, getPaginatedTasksByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPaginatedTasksByUserIDRow{}
	for rows.Next() {
		var i GetPaginatedTasksByUserIDRow
		if err := rows.Scan(
			&i.TableID,
			&i.UserID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.TableID_2,
			&i.TaskID,
			&i.UserID_2,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTask = `-- name: GetTask :one
SELECT table_id, task_id, user_id, data_range, reminder_period, created_at
FROM tasks
WHERE task_id = $1
LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, taskID uuid.UUID) (Tasks, error) {
	row := q.queryRow(ctx, q.getTaskStmt, getTask, taskID)
	var i Tasks
	err := row.Scan(
		&i.TableID,
		&i.TaskID,
		&i.UserID,
		&i.DataRange,
		&i.ReminderPeriod,
		&i.CreatedAt,
	)
	return i, err
}

const getTasksByUserEmail = `-- name: GetTasksByUserEmail :many
SELECT u.table_id, u.user_id, email, first_name, last_name, u.created_at, t.table_id, task_id, t.user_id, data_range, reminder_period, t.created_at
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.email = &1
`

type GetTasksByUserEmailRow struct {
	TableID        sql.NullInt64 `json:"tableID"`
	UserID         uuid.UUID     `json:"userID"`
	Email          string        `json:"email"`
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	CreatedAt      sql.NullTime  `json:"createdAt"`
	TableID_2      sql.NullInt64 `json:"tableID2"`
	TaskID         uuid.UUID     `json:"taskID"`
	UserID_2       uuid.UUID     `json:"userID2"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
	CreatedAt_2    sql.NullTime  `json:"createdAt2"`
}

func (q *Queries) GetTasksByUserEmail(ctx context.Context) ([]GetTasksByUserEmailRow, error) {
	rows, err := q.query(ctx, q.getTasksByUserEmailStmt, getTasksByUserEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTasksByUserEmailRow{}
	for rows.Next() {
		var i GetTasksByUserEmailRow
		if err := rows.Scan(
			&i.TableID,
			&i.UserID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.TableID_2,
			&i.TaskID,
			&i.UserID_2,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTasksByUserID = `-- name: GetTasksByUserID :many
SELECT u.table_id, u.user_id, email, first_name, last_name, u.created_at, t.table_id, task_id, t.user_id, data_range, reminder_period, t.created_at
FROM users u
         INNER JOIN tasks t
                    ON t.user_id = u.user_id
WHERE u.user_id = &1
`

type GetTasksByUserIDRow struct {
	TableID        sql.NullInt64 `json:"tableID"`
	UserID         uuid.UUID     `json:"userID"`
	Email          string        `json:"email"`
	FirstName      string        `json:"firstName"`
	LastName       string        `json:"lastName"`
	CreatedAt      sql.NullTime  `json:"createdAt"`
	TableID_2      sql.NullInt64 `json:"tableID2"`
	TaskID         uuid.UUID     `json:"taskID"`
	UserID_2       uuid.UUID     `json:"userID2"`
	DataRange      interface{}   `json:"dataRange"`
	ReminderPeriod sql.NullInt64 `json:"reminderPeriod"`
	CreatedAt_2    sql.NullTime  `json:"createdAt2"`
}

func (q *Queries) GetTasksByUserID(ctx context.Context) ([]GetTasksByUserIDRow, error) {
	rows, err := q.query(ctx, q.getTasksByUserIDStmt, getTasksByUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTasksByUserIDRow{}
	for rows.Next() {
		var i GetTasksByUserIDRow
		if err := rows.Scan(
			&i.TableID,
			&i.UserID,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.TableID_2,
			&i.TaskID,
			&i.UserID_2,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaginatedTasks = `-- name: ListPaginatedTasks :many
SELECT table_id, task_id, user_id, data_range, reminder_period, created_at
FROM tasks
ORDER BY table_id
LIMIT $1 OFFSET $2
`

type ListPaginatedTasksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPaginatedTasks(ctx context.Context, arg ListPaginatedTasksParams) ([]Tasks, error) {
	rows, err := q.query(ctx, q.listPaginatedTasksStmt, listPaginatedTasks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tasks{}
	for rows.Next() {
		var i Tasks
		if err := rows.Scan(
			&i.TableID,
			&i.TaskID,
			&i.UserID,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasks = `-- name: ListTasks :many
SELECT table_id, task_id, user_id, data_range, reminder_period, created_at
FROM tasks
ORDER BY time
`

func (q *Queries) ListTasks(ctx context.Context) ([]Tasks, error) {
	rows, err := q.query(ctx, q.listTasksStmt, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tasks{}
	for rows.Next() {
		var i Tasks
		if err := rows.Scan(
			&i.TableID,
			&i.TaskID,
			&i.UserID,
			&i.DataRange,
			&i.ReminderPeriod,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
