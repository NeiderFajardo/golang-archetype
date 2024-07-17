package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/NeiderFajardo/internal/products/domain"
	"github.com/NeiderFajardo/pkg/apierrors"
	"github.com/NeiderFajardo/pkg/database"
	"github.com/NeiderFajardo/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	dbClient *database.MongoDatabase
}

func NewProductRepository(params ProductRepositoryParams) domain.IProductRepository {
	return &ProductRepository{
		dbClient: params.DbClient,
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
			logger.Info(err.Error())
			return nil, errors.New("timeout getting product")
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
			return 0, errors.New("timeout creating product")
		}
		if mongo.IsDuplicateKeyError(err) {
			return 0, apierrors.BadRequest("Product already exists", "conflict", "id")
		}
		return 0, err
	}
	return product.ID, nil
}

func (pr *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: product.ID}}
	update := bson.D{{Key: "$set", Value: product}}
	_, err := pr.dbClient.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if mongo.IsTimeout(err) {
			return errors.New("timeout updating product")
		}
		return err
	}
	return nil
}
