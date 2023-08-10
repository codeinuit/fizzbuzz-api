package mock

import (
	"errors"

	"github.com/codeinuit/fizzbuzz-api/pkg/models"
)

type DatabaseMock struct {
	awaited any
}

func NewDatabaseMock() *DatabaseMock {
	return &DatabaseMock{}
}

func (db DatabaseMock) CountUsage() (models.Stats, error) {
	awaited, ok := db.awaited.(models.Stats)
	if !ok {
		return models.Stats{}, errors.New("error")
	}

	return awaited, nil
}

func (db DatabaseMock) UsageUpdate(m models.Stats) {

}
