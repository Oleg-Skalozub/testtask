package repository

import (
	"github.com/Oleg-Skalozub/testtask/src/domain/entity"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/db"
)

// DataRepository ...
type DataRepository struct {
	db db.SqlBD
}

// NewDataRepository ...
func NewDataRepository() *DataRepository {
	return &DataRepository{
		db: db.SQLBD,
	}
}

// GetData ...
func (dr DataRepository) GetData(day, month int) ([]entity.DataResponse, error) {
	var data []entity.DataResponse
	dr.db.Table(entity.TableName).Select("event_type, count(year) as result").Where("day = ? AND month=?", day, month).Group("event_type").Find(&data)
	return data, nil
}

// SaveData ...
func (dr DataRepository) SaveData(day, month, eventType int, year, title string) error {
	return dr.db.Save(entity.DataDB{
		Day:       day,
		Month:     month,
		EventType: eventType,
		Year:      year,
		Title:     title,
	})
}
