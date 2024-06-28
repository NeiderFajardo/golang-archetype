package infrastructure

import (
	"context"
	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/database"
	"testing"
)

func TestProductRepository_GetByID_NotFound(t *testing.T) {

	testDatabase := database.SetupTestDatabase()
	defer testDatabase.TearDown()

	productRepository := NewProductRepository(testDatabase.MongoDatabase)
	_, err := productRepository.GetByID(context.Background(), 1)
	resultError, ok := err.(*apierrors.ApiError)
	if !ok {
		t.Errorf("Error not_found expected, got %s", err)
	}
	if resultError.Code != "not_found" {
		t.Errorf("Error not_found expected, got %s", resultError.Code)
	}
}
