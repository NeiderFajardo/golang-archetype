package database

import (
	"context"
	"fmt"
	"github.com/NeiderFajardo/config"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"
)

type TestDatabase struct {
	MongoDatabase *MongoDatabase
	container     testcontainers.Container
}

func createMongoContainer(ctx context.Context) (testcontainers.Container, *MongoDatabase, error) {
	var env = map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": "root",
		"MONGO_INITDB_ROOT_PASSWORD": "password",
		"MONGO_INITDB_DATABASE":      "test",
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{"27017/tcp"},
			Env:          env,
			WaitingFor:   wait.ForHTTP("/").WithPort("27017"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return container, nil, fmt.Errorf("failed to get container external port: %v", err)
	}

	log.Println("mongo container ready and running at port: ", p.Port())

	uri := fmt.Sprintf("mongodb://root:password@localhost:%s", p.Port())
	dbConfig := config.LoadMongoConfig(uri, "test", "products")
	client := NewMongoClient(dbConfig)
	if err != nil {
		return container, client, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, client, nil
}

func SetupTestDatabase() *TestDatabase {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	container, dbInstance, err := createMongoContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup test ", err)
	}

	return &TestDatabase{
		container:     container,
		MongoDatabase: dbInstance,
	}
}

func (tdb *TestDatabase) TearDown() {
	_ = tdb.container.Terminate(context.Background())
}
