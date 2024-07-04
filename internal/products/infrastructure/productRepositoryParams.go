package infrastructure

import (
	"github.com/NeiderFajardo/pkg/database"
)

type ProductRepositoryParams struct {
	// fx.In

	DbClient *database.MongoDatabase
}
