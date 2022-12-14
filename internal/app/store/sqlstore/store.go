package sqlstore

import (
	"database/sql"

	"github.com/MixwelInc/http-rest-api/internal/app/store"
	_ "github.com/lib/pq" // anon for not using it methods in code
)

// Store ...
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
