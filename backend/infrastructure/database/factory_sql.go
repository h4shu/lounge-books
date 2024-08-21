package database

import (
	"errors"

	"github.com/h4shu/lounge-books/application/repositories"
)

const (
	InstancePostgres int = iota
)

var (
	errInvalidSQLDatabaseInstance = errors.New("invalid sql db instance")
)

func NewDatabaseSQLFactory(instance int) (repositories.SQL, error) {
	switch instance {
	case InstancePostgres:
		return newPostgresHandler(newConfigPostgres())
	default:
		return nil, errInvalidSQLDatabaseInstance
	}
}
