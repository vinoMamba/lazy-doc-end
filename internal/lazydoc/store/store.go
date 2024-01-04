package store

import (
	"sync"

	"gorm.io/gorm"
)

type IStore interface {
	Users() UserStore
}

type dataStore struct {
	db *gorm.DB
}

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
