package user

import "database/sql"

// Store provides all functions to execute db queries
type Store interface {
	Querier
}

type userStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a user store object returning the store interface.
func NewStore(db *sql.DB) Store {
	us := &userStore{
		Queries: New(db),
		db:      db,
	}

	return us
}
