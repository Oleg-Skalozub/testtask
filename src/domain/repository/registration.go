package repository

import "github.com/Oleg-Skalozub/testtask/src/infrastructure/db/repository"

// NewDataRepository ...
func NewDataRepository() DataRepository {
	return repository.NewDataRepository()
}
