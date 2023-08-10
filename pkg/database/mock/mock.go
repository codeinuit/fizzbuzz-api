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

func (db *DatabaseMock) UsageUpdate(m models.Stats) {
	db.awaited = models.Stats{
		Int1:    m.Int1,
		Int2:    m.Int2,
		Int3:    m.Int3,
		String1: m.String1,
		String2: m.String2,
		Use:     m.Use + 1,
	}
}
