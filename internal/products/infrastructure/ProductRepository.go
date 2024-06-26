package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NeiderFajardo/internal/products/domain"
	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	dbClient *database.MongoDatabase
}

func NewProductRepository(dbClient *database.MongoDatabase) domain.IProductRepository {
	return &ProductRepository{
		dbClient: dbClient,
	}
}

func (pr *ProductRepository) GetByID(ctx context.Context, id int) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	product := domain.Product{}
	err := pr.dbClient.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Product not found: %d", id)
			return nil, apierrors.NotFound(errorMessage, "not_found", "id")
		}
		if mongo.IsTimeout(err) {
			return nil, errors.New("Timeout getting product")
		}
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) Create(ctx context.Context, product *domain.Product) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	_, err := pr.dbClient.Collection.InsertOne(ctx, product)
	if err != nil {
		if mongo.IsTimeout(err) {
			return 0, errors.New("Timeout creating product")
		}
		if mongo.IsDuplicateKeyError(err) {
			return 0, apierrors.BadRequest("Product already exists", "conflict", "id")
		}
		return 0, err
	}
	return product.ID, nil
}
