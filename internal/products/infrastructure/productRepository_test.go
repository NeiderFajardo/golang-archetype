package infrastructure

import (
	"context"
	"testing"

	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/database"
)

func TestProductRepository_GetByID_NotFound(t *testing.T) {

	testDatabase := database.SetupTestDatabase()
	params := ProductRepositoryParams{
		DbClient: testDatabase.MongoDatabase,
	}
	defer testDatabase.TearDown()

	productRepository := NewProductRepository(params)
	_, err := productRepository.GetByID(context.Background(), 1)
	resultError, ok := err.(*apierrors.ApiError)
	if !ok {
		t.Errorf("Error not_found expected, got %s", err)
	}
	if resultError.Code != "not_found" {
		t.Errorf("Error not_found expected, got %s", resultError.Code)
	}
}
