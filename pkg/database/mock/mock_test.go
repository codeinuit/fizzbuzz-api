package mock_test

import (
	"testing"

	"github.com/codeinuit/fizzbuzz-api/pkg/database"
	"github.com/codeinuit/fizzbuzz-api/pkg/database/mock"
)

func TestMockImplementation(t *testing.T) {
	var db database.Database
	db = mock.NewDatabaseMock()
	_ = db
}
