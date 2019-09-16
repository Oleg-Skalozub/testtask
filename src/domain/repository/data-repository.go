package repository

import "github.com/Oleg-Skalozub/testtask/src/domain/entity"

// DataRepository
type DataRepository interface {
	GetData(day, month int) ([]entity.DataResponse, error)
	SaveData(day, month, eventType int, year, title string) error
}
