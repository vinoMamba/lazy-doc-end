package store

import (
	"sync"

	"gorm.io/gorm"
)

type IStore interface {
	Users() UserStore
	Projects() ProjectStore
}

type dataStore struct {
	db *gorm.DB
}

var _ IStore = (*dataStore)(nil)

var (
	once sync.Once
	Ds   *dataStore
)

func NewStore(db *gorm.DB) *dataStore {
	once.Do(func() {
		Ds = &dataStore{db}
	})
	return Ds
}

func (s *dataStore) Users() UserStore {
	return newUsers(s.db)
}

func (s *dataStore) Projects() ProjectStore {
	return newProjects(s.db)
}
