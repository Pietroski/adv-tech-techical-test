package task

import "database/sql"

// Store provides all functions to execute db queries
type Store interface {
	Querier
}

type taskStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a task store object returning the store interface.
func NewStore(db *sql.DB) Store {
	ts := &taskStore{
		Queries: New(db),
		db:      db,
	}

	return ts
}
