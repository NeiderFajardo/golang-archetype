package infrastructure

import (
	"context"
	"errors"

	"github.com/NeiderFajardo/internal/products/domain"
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

func (pr *ProductRepository) GetByID(id int) (*domain.Product, error) {
	ctx := context.Background()
	filter := bson.D{{Key: "id", Value: id}}
	product := domain.Product{}
	err := pr.dbClient.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (pr *ProductRepository) Create(product *domain.Product) (int, error) {
	ctx := context.Background()
	_, err := pr.dbClient.Collection.InsertOne(ctx, product)
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}
